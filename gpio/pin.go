package gpio

type Pin struct {
	gpio *gpioSingleton
	pin  uint8
}

func NewPin(gpio *gpioSingleton, pin uint8) *Pin {
	this := &Pin{
		gpio: gpio,
		pin:  pin,
	}
	return this
}

// Get the GPIO number for this pin.
func (this *Pin) Pin() uint8 {
	return this.pin
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
	this.gpio.Pull(this, pull)
}

// Set pin Direction
func (this *Pin) Mode(dir Direction) {
	this.gpio.Mode(this, dir)
}

// Read pin state (high/low)
func (this *Pin) Read() State {
	return this.gpio.Read(this)
}

// Set pin state (high/low)
func (this *Pin) Write(state State) {
	this.gpio.Write(this, state)
}

// Takes a range from 0% to 100% as an integer and sets the pin.High() to pulse at that percentage of 1/500 of a second.
func (this *Pin) Modulate(modulation int) {
	this.gpio.Modulate(this, modulation)
}
