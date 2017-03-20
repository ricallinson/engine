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

func (this *LED) Pin() int {
	return int(this.pin)
}

func (this *LED) Toggle() {
	this.pin.Toggle()
	this.log()
}

func (this *LED) Set(val float32) {
	if val <= 0 {
		this.pin.Low()
	} else {
		this.pin.High()
	}
	this.log()
}

func (this *LED) log() {
	log.Print("LED on pin ", this.pin, " was set to ", this.pin.Read())
}
