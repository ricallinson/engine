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

// Global list of used pins.
var pinsUsed = make([]bool, 25, 25)

// Global engine in use flag.
var locked bool

type Engine struct {
	testing bool
}

//
func Start(mock bool) *Engine {
	if locked {
		log.Panic("The engine is already being used.")
	}
	rpio.Mock = mock
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		log.Panic("Could not access GPIO memory.", err)
	}
	this := &Engine{}
	locked = true
	return this
}

func (this *Engine) Stop() {
	pinsUsed = make([]bool, 26, 26)
	locked = false
	rpio.Close()
}

func (this *Engine) NewLED(pin int) *LED {
	this.registerPin(pin)
	return NewLED(pin)
}

func (this *Engine) NewMotor(pinA int, pinB int, pinE int, reversed bool) *Motor {
	this.registerPin(pinA)
	this.registerPin(pinB)
	this.registerPin(pinE)
	return NewMotor(pinA, pinB, pinE, reversed)
}

func (this *Engine) NewIRSensor(pin int) *IRSensor {
	this.registerPin(pin)
	return NewIRSensor(pin)
}

// If a pin has already been used then it results in a fatal error.
func (this *Engine) registerPin(pin int) {
	if pin >= len(pinsUsed) || pin < 1 {
		log.Panic("Pin number ", pin, " is out of range.")
	}
	if pinsUsed[pin] {
		log.Panic("Pin number ", pin, " has already been used.")
	}
	pinsUsed[pin] = true
}
