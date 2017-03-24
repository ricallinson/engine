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
	//...
}

// Set pin Direction
func (this *Pin) Mode(dir Direction) {
	//...
}

// Set pin state (high/low)
func (this *Pin) Write(state State) {
	//...
}

// Takes a range from 0% to 100% as an integer and sets the pin.High() to pulse at that percentage of 1/500 of a second.
func (this *Pin) WriteModulation(modulation int) {
	//...
}

// Read pin state (high/low)
func (this *Pin) Read() State {
	return Low
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
	//...
}
