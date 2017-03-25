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

type PinInput struct {
	*Pin
	name string
}

// Returns a new instance of PinInput.
func NewPinInput(pin *gpio.GpioPin) *PinInput {
	this := &PinInput{
		Pin:  NewPin(pin),
		name: "PinInput",
	}
	this.Input()
	return this
}

// Returns the current value of this instances PinInput.
// The range is 0-1 rounded up where 0 is obstacle detected and 1 is no obstacle detected.
func (this *PinInput) Get() float32 {
	val := float32(this.Read())
	this.log(val)
	return val
}

// Logs state of the assigned pin.
func (this *PinInput) log(val float32) {
	log.Print(this.name, " on pin ", this.PinOut(), " read a value of ", val)
}
