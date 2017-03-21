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

type Button struct {
	pin rpio.Pin
}

// Returns a new instance of Button.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
func NewButton(pin int) *Button {
	this := &Button{
		pin: rpio.Pin(pin),
	}
	this.pin.Input()
	return this
}

// Returns the pin that this instance is controlled by.
func (this *Button) Pin() int {
	return int(this.pin)
}

// Returns the current value of this instances Button.
// The range is 0-1 rounded up where 0 is obstacle detected and 1 is no obstacle detected.
func (this *Button) Get() float32 {
	val := this.pin.Read()
	this.log(val)
	return float32(val)
}

// Logs state of the assigned pin.
func (this *Button) log(val rpio.State) {
	log.Print("Button on pin ", this.pin, " read a value of ", val)
}
