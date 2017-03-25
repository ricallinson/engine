package gpio

import (
	"time"
)

type Pin struct {
	gpio            *gpioSingleton
	pin             uint8
	modulation      int
	modulationPhase chan int
}

func NewPin(gpio *gpioSingleton, pin uint8) *Pin {
	this := &Pin{
		gpio:            gpio,
		pin:             pin,
		modulationPhase: make(chan int),
	}
	return this
}

// Get the GPIO number for this pin.
func (this *Pin) Pin() uint8 {
	return this.pin
}

// Close the modulationPhase channel.
func (this *Pin) Close() {
	close(this.modulationPhase)
}

// Set pin as Input
func (this *Pin) Input() {
	this.Mode(Input)
}

// Set pin as Output
func (this *Pin) Output() {
	this.Mode(Output)
}

// Set pin High
func (this *Pin) High() {
	this.Write(High)
}

// Set pin Low
func (this *Pin) Low() {
	this.Write(Low)
}

// Toggle pin state
func (this *Pin) Toggle() {
	if this.Read() == High {
		this.Low()
	} else {
		this.High()
	}
}

// Pull up pin
func (this *Pin) PullUp() {
	this.Pull(PullUp)
}

// Pull down pin
func (this *Pin) PullDown() {
	this.Pull(PullDown)
}

// Disable pullup/down on pin
func (this *Pin) PullOff() {
	this.Pull(PullOff)
}

// Set a given pull up/down mode
func (this *Pin) Pull(pull Pull) {
	this.gpio.Pull(this.Pin(), pull)
}

// Set pin Direction
func (this *Pin) Mode(dir Direction) {
	this.gpio.Mode(this.Pin(), dir)
}

// Read pin state (high/low)
func (this *Pin) Read() State {
	return this.gpio.Read(this.Pin())
}

// Set pin state (high/low)
func (this *Pin) Write(state State) {
	this.gpio.Write(this.Pin(), state)
}

// Takes a range from 0% to 100% as an integer and sets the pin.High() to pulse at that percentage of 1/500 of a second.
func (this *Pin) Modulate(modulation int) {
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
	go this.modulatePin()
}

// Software implemented Pulse Width Modulation (PWM) at every 2ms.
// A channel is used to stop the routine when Close() is called.
func (this *Pin) modulatePin() {
	// Create the int to store the High microsecond time.
	var high int
	// Started modulating on pin.
	// Check that modulation value is in range.
	for this.modulation > 0 && this.modulation < 100 {
		switch <-this.modulationPhase {
		case 0:
			return
		case 1:
			// The modulation is 200uS so multiply the percentage by 2.
			high = this.modulation * 2
			this.High()
			// Sleep for pulse high duration.
			time.Sleep(time.Duration(high) * time.Microsecond)
			this.modulationPhase <- 2
		case 2:
			this.Low()
			// Sleep for pulse low duration. The remainder of the 200uS used by pulse High.
			time.Sleep(time.Duration(200-high) * time.Microsecond)
			// Get the current PWM value for this pin.
			this.modulationPhase <- 1
		}
	}
}
