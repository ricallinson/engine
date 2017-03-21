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

type LED struct {
	pin rpio.Pin
}

// Returns a new instance of LED.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
// Controls a light admitting diode (LED).
func NewLED(pin int) *LED {
	this := &LED{
		pin: rpio.Pin(pin),
	}
	this.pin.Output()
	return this
}

// Returns the pin that this instance is controlled by.
func (this *LED) Pin() int {
	return int(this.pin)
}

func (this *LED) On() {
	this.Set(1)
}

func (this *LED) Off() {
	this.Set(0)
}

func (this *LED) Toggle() {
	this.pin.Toggle()
	this.log()
}

// Set the current value of this instances LED.
// The range is 0-1 rounded up where 0 is off and 1 is on.
func (this *LED) Set(val float32) {
	if val > 0 {
		this.pin.High()
	} else {
		this.pin.Low()
	}
	this.log()
}

// Logs state of the assigned pin.
func (this *LED) log() {
	log.Print("LED on pin ", this.pin, " was set to ", this.pin.Read())
}
