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

type PinOutput struct {
	pin  rpio.Pin
	name string
}

// Returns a new instance of PinOutput.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
func NewPinOutput(pin int) *PinOutput {
	this := &PinOutput{
		pin:  rpio.Pin(pin),
		name: "PinOutput",
	}
	this.pin.Output()
	return this
}

// Returns the name of this instance.
func (this *PinOutput) String() string {
	return this.name
}

// Returns the pin that this instance is controlled by.
func (this *PinOutput) Pin() int {
	return int(this.pin)
}

// Set the current value of this instances PinOutput.
// The range is 0-1 rounded up where 0 is off and 1 is on.
func (this *PinOutput) Set(val float32) {
	if val > 0 {
		this.pin.High()
	} else {
		this.pin.Low()
	}
	this.log()
}

// Logs state of the assigned pin.
func (this *PinOutput) log() {
	log.Print(this.name, " on pin ", this.pin, " read a value of ", this.pin.Read())
}
