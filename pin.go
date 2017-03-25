//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package engine

import (
	"github.com/ricallinson/engine/gpio"
)

type Pin struct {
	*gpio.GpioPin
	name string
}

// Returns a new instance of Pin.
func NewPin(pin *gpio.GpioPin) *Pin {
	this := &Pin{
		GpioPin: pin,
		name:    "Pin",
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
