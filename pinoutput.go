//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package engine

import (
	"log"
)

type PinOutput struct {
	*Pin
	name string
}

// Returns a new instance of PinOutput.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
func NewPinOutput(pin int) *PinOutput {
	this := &PinOutput{
		Pin:  NewPin(pin),
		name: "PinOutput",
	}
	this.pin.Output()
	return this
}

// Set the current value of this instances PinOutput.
// The range is 0-1 rounded up where 0 is off and 1 is on.
func (this *PinOutput) Set(val float32) {
	this.pin.WritePWM(int(val * 100))
	this.log(val)
}

// Logs state of the assigned pin.
func (this *PinOutput) log(val float32) {
	log.Print(this.name, " on pin ", this.pin, " read a value of ", val)
}
