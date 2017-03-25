//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package gpio

import (
	"time"
)

type GpioPin struct {
	gpio       *GpioSingleton
	pin        uint8
	modulation int
	direction  Direction
	pull       Pull
	lastWrite  State
}

func NewPin(gpio *GpioSingleton, pin uint8) *GpioPin {
	this := &GpioPin{
		gpio: gpio,
		pin:  pin,
	}
	return this
}

// Get the GPIO number for this pin.
func (this *GpioPin) PinOut() uint8 {
	return this.pin
}

// Set pin as Input.
func (this *GpioPin) Input() {
	this.Mode(Input)
}

// Set pin as Output.
func (this *GpioPin) Output() {
	this.Mode(Output)
}

// Set pin High.
func (this *GpioPin) High() {
	this.Write(High)
}

// Set pin Low.
func (this *GpioPin) Low() {
	this.Write(Low)
}

// Toggle pin state.
func (this *GpioPin) Toggle() {
	if this.Read() == High {
		this.Low()
	} else {
		this.High()
	}
}

// Pull up pin.
func (this *GpioPin) PullUp() {
	this.Pull(PullUp)
}

// Pull down pin.
func (this *GpioPin) PullDown() {
	this.Pull(PullDown)
}

// Disable pullup/down on pin.
func (this *GpioPin) PullOff() {
	this.Pull(PullOff)
}

// Set a given pull up/down mode.
func (this *GpioPin) Pull(pull Pull) {
	this.pull = pull
	this.gpio.Pull(this.PinOut(), pull)
}

// Set a given pull up/down mode.
func (this *GpioPin) GetPull() Pull {
	return this.pull
}

// Set pin Direction.
func (this *GpioPin) Mode(dir Direction) {
	this.direction = dir
	this.gpio.Mode(this.PinOut(), dir)
}

// Get pin Direction.
func (this *GpioPin) GetMode() Direction {
	return this.direction
}

// Read pin state (high/low).
func (this *GpioPin) Read() State {
	return this.gpio.Read(this.PinOut())
}

// Set pin state (high/low).
func (this *GpioPin) Write(state State) {
	this.lastWrite = state
	this.gpio.Write(this.PinOut(), state)
}

// Set pin state (high/low).
func (this *GpioPin) LastWrite() State {
	return this.lastWrite
}

// Takes a range from 0% to 100% as an integer and sets the pin.High() to pulse at that percentage of 1/500 of a second.
func (this *GpioPin) Modulate(modulation int) {
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
	go this.modulateGpioPin()
}

// Get pin modulation.
func (this *GpioPin) GetModulation() int {
	return this.modulation
}

// Software implemented Pulse Width Modulation (PWM) at every 2ms.
func (this *GpioPin) modulateGpioPin() {
	width := 100 // in uS - this should be 200uS but in testing something is off.
	widthHighOffset := 50
	var high int
	var low int
	// Check that modulation value is in range.
	for this.modulation > 0 && this.modulation < 100 {
		// Set the high and low modulation timing to fit in the width.
		high = this.modulation * (width / 100)
		low = width - high
		this.High()
		// Sleep for pulse high duration.
		time.Sleep(time.Duration(high-widthHighOffset) * time.Microsecond)
		this.Low()
		// Sleep for pulse low duration.
		time.Sleep(time.Duration(low) * time.Microsecond)
	}
}
