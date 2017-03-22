//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package engine

import (
	"github.com/ricallinson/engine/rpio"
)

type Pin struct {
	pin  rpio.Pin
	name string
}

// Returns a new instance of Pin.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
func NewPin(pin int) *Pin {
	this := &Pin{
		pin:  rpio.Pin(pin),
		name: "Pin",
	}
	return this
}

// Sets the name of this instance.
func (this *Pin) Name(name string) {
	this.name = name
}

// Returns the name of this instance.
func (this *Pin) String() string {
	return this.name
}

// Returns the pin that this instance is controlled by.
func (this *Pin) PinNumber() int {
	return int(this.pin)
}
