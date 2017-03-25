//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package engine

import (
	"github.com/ricallinson/engine/gpio"
)

type IRSensor struct {
	*PinInput
}

// Returns a new instance of IRSensor.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
// Controls an infra red sensor array with VCC, GND and DO connectors.
func NewIRSensor(pin *gpio.GpioPin) *IRSensor {
	this := &IRSensor{
		NewPinInput(pin),
	}
	return this
}
