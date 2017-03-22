//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package engine

type LED struct {
	*PinOutput
}

// Returns a new instance of LED.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
// Controls a light admitting diode (LED).
func NewLED(pin int) *LED {
	this := &LED{
		NewPinOutput(pin),
	}
	return this
}

// Turn the LED on.
func (this *LED) On() {
	this.Set(1)
}

// Turn the LED off.
func (this *LED) Off() {
	this.Set(0)
}

// Toggle the current state of the LED.
func (this *LED) Toggle() {
	this.pin.Toggle()
	this.log()
}
