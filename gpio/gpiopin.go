//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package gpio

import (
	"time"
	"fmt"
)

type GpioPin struct {
	gpio      *GpioSingleton
	pin       uint8
	dutyCycle int
	hertz     int
	direction Direction
	pull      Pull
	lastWrite State
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

// Takes duty cycle from 0% to 100% and hertz less than 100hz.
func (this *GpioPin) Modulate(dutyCycle int, hertz int) {
	// Frequency of a 100Hz which is 100 times a second (based on https://projects.drogon.net/raspberry-pi/wiringpi/software-pwm-library/).
	if hertz > 100 || hertz < 0 {
		this.hertz = 100
	} else {
		this.hertz = hertz
	}
	// If dutyCycle is 0 or less then reset stored dutyCycle and call pin.Low().
	if dutyCycle < 1 {
		this.dutyCycle = 0
		this.Low()
		return
	}
	// If dutyCycle is 100 or greater then reset stored dutyCycle and call pin.High().
	if dutyCycle > 99 {
		this.dutyCycle = 0
		this.High()
		return
	}
	// If there is already a dutyCycle value then update it and return.
	if this.dutyCycle > 0 {
		this.dutyCycle = dutyCycle
		return
	}
	// If none of the above are true then store the dutyCycle percentage for the pin.
	this.dutyCycle = dutyCycle
	// Start the dutyCycle routine that will run until the dutyCycle value is out of range.
	go this.modulateGpioPin()
}

// Get the pins dutyCycle.
func (this *GpioPin) GetModulation() int {
	return int(this.dutyCycle)
}

// Software implemented Pulse Width Modulation (PWM) at 100Hz.
//  - this should be 200uS but in testing a value of 100uS ranges from 210uS to 300uS.
func (this *GpioPin) modulateGpioPin() {
	width := 1000000 / this.hertz // in microseconds.
	var high int
	var low int
	// Check that dutyCycle value is in range.
	for this.dutyCycle > 0 && this.dutyCycle < 100 {
		// Set the high and low dutyCycle timing to fit in the width.
		high = this.dutyCycle * (width / 100)
		low = width - high
		fmt.Println(high, low, high + low, width)
		this.High()
		// Sleep for pulse high duration.
		time.Sleep(time.Duration(high) * time.Microsecond)
		this.Low()
		// Sleep for pulse low duration.
		time.Sleep(time.Duration(low) * time.Microsecond)
	}
}
