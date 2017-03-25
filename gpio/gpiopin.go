package gpio

import (
	"time"
)

type gpioPin struct {
	gpio       *gpioSingleton
	pin        uint8
	modulation int
	direction  Direction
	pull       Pull
	lastWrite  State
}

func NewPin(gpio *gpioSingleton, pin uint8) *gpioPin {
	this := &gpioPin{
		gpio: gpio,
		pin:  pin,
	}
	return this
}

// Get the GPIO number for this pin.
func (this *gpioPin) Pin() uint8 {
	return this.pin
}

// Set pin as Input
func (this *gpioPin) Input() {
	this.Mode(Input)
}

// Set pin as Output
func (this *gpioPin) Output() {
	this.Mode(Output)
}

// Set pin High
func (this *gpioPin) High() {
	this.Write(High)
}

// Set pin Low
func (this *gpioPin) Low() {
	this.Write(Low)
}

// Toggle pin state
func (this *gpioPin) Toggle() {
	if this.Read() == High {
		this.Low()
	} else {
		this.High()
	}
}

// Pull up pin
func (this *gpioPin) PullUp() {
	this.Pull(PullUp)
}

// Pull down pin
func (this *gpioPin) PullDown() {
	this.Pull(PullDown)
}

// Disable pullup/down on pin
func (this *gpioPin) PullOff() {
	this.Pull(PullOff)
}

// Set a given pull up/down mode
func (this *gpioPin) Pull(pull Pull) {
	this.pull = pull
	this.gpio.Pull(this.Pin(), pull)
}

// Set a given pull up/down mode
func (this *gpioPin) GetPull() Pull {
	return this.pull
}

// Set pin Direction.
func (this *gpioPin) Mode(dir Direction) {
	this.direction = dir
	this.gpio.Mode(this.Pin(), dir)
}

// Get pin Direction.
func (this *gpioPin) GetMode() Direction {
	return this.direction
}

// Read pin state (high/low)
func (this *gpioPin) Read() State {
	return this.gpio.Read(this.Pin())
}

// Set pin state (high/low)
func (this *gpioPin) Write(state State) {
	this.lastWrite = state
	this.gpio.Write(this.Pin(), state)
}

// Set pin state (high/low)
func (this *gpioPin) LastWrite() State {
	return this.lastWrite
}

// Takes a range from 0% to 100% as an integer and sets the pin.High() to pulse at that percentage of 1/500 of a second.
func (this *gpioPin) Modulate(modulation int) {
	// If modulation is 0 or less then reset stored modulation and call pin.Low().
	if modulation < 1 {
		this.modulation = 0
		this.Low()
		return
	}
	// If modulation is 100 or greater then reset stored modulation and call pin.High().
	if modulation > 99 {
		this.modulation = 0
		this.High()
		return
	}
	// If there is already a modulation value then update it and return.
	if this.modulation > 0 {
		this.modulation = modulation
		return
	}
	// If none of the above are true then store the modulation percentage for the pin.
	this.modulation = modulation
	// Start the modulation routine that will run until the modulation value is out of range.
	go this.modulategpioPin()
}

// Get pin modulation.
func (this *gpioPin) GetModulation() int {
	return this.modulation
}

// Software implemented Pulse Width Modulation (PWM) at every 2ms.
// A channel is used to stop the routine when Close() is called.
func (this *gpioPin) modulategpioPin() {
	// Create the int to store the High microsecond time.
	var high int
	var phase State
	// Check that modulation value is in range.
	for this.modulation > 0 && this.modulation < 100 {
		switch phase {
		case High:
			// The modulation is 200uS so multiply the percentage by 2.
			high = this.modulation * 2
			this.High()
			// Sleep for pulse high duration.
			time.Sleep(time.Duration(high) * time.Microsecond)
			phase = Low
		case Low:
			this.Low()
			// Sleep for pulse low duration. The remainder of the 200uS used by pulse High.
			time.Sleep(time.Duration(200-high) * time.Microsecond)
			phase = High
		}
	}
}
