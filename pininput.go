//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package engine

import (
	"github.com/ricallinson/engine/rpio"
	"log"
)

type PinInput struct {
	*Pin
	name string
}

// Returns a new instance of PinInput.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
func NewPinInput(pin int) *PinInput {
	this := &PinInput{
		Pin:  NewPin(pin),
		name: "PinInput",
	}
	this.pin.Input()
	return this
}

// Returns the current value of this instances PinInput.
// The range is 0-1 rounded up where 0 is obstacle detected and 1 is no obstacle detected.
func (this *PinInput) Get() float32 {
	val := this.pin.Read()
	this.log(val)
	return float32(val)
}

// Logs state of the assigned pin.
func (this *PinInput) log(val rpio.State) {
	log.Print(this.name, " on pin ", this.pin, " read a value of ", val)
}
