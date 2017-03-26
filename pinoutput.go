//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package engine

import (
	"github.com/ricallinson/engine/gpio"
	"log"
)

type PinOutput struct {
	*Pin
	name string
}

// Returns a new instance of PinOutput.
func NewPinOutput(pin *gpio.GpioPin) *PinOutput {
	this := &PinOutput{
		Pin:  NewPin(pin),
		name: "PinOutput",
	}
	this.Output()
	return this
}

// Set the current value of this instances PinOutput.
// The range is 0-1 rounded up where 0 is off and 1 is on.
func (this *PinOutput) Set(val float32) {
	this.Modulate(int(val * 100), 100)
	this.log(val)
}

// Logs state of the assigned pin.
func (this *PinOutput) log(val float32) {
	log.Print(this.name, " on pin ", this.PinOut(), " read a value of ", val)
}
