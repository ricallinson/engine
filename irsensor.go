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

type IRSensor struct {
	pin rpio.Pin
}

// Returns a new instance of IRSensor.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
// Controls an infra red sensor array with VCC, GND and DO connectors.
func NewIRSensor(pin int) *IRSensor {
	this := &IRSensor{
		pin: rpio.Pin(pin),
	}
	this.pin.Input()
	return this
}

// Returns the pin that this instance is controlled by.
func (this *IRSensor) Pin() int {
	return int(this.pin)
}

// Returns the current value of this instances IRSensor.
// The range is 0-1 rounded up where 0 is obstacle detected and 1 is no obstacle detected.
func (this *IRSensor) Get() float32 {
	val := this.pin.Read()
	this.log(val)
	return float32(val)
}

// Logs state of the assigned pin.
func (this *IRSensor) log(val rpio.State) {
	log.Print("IRSensor on pin ", this.pin, " read a value of ", val)
}
