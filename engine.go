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

// Interface for managing hardware pins and instantiating control objects such as Motors, LEDs and IRSensors.
type Engine struct {
	testing bool
}

// Starts the engine returning an instance of Engine.
// Subsequent calls to Start() without calling Stop() first will generate a panic.
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

// Stops the engine and releases all resources.
// It must be called before Start() can be called again.
func (this *Engine) Stop() {
	pinsUsed = make([]bool, 26, 26)
	locked = false
	rpio.Close()
}

// Returns a new instance of Switch.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
func (this *Engine) NewSwitch(pin int) *Switch {
	this.registerPin(pin)
	return NewSwitch(pin)
}

// Returns a new instance of LED.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
func (this *Engine) NewLED(pin int) *LED {
	this.registerPin(pin)
	return NewLED(pin)
}

// Returns a new instance of Motor.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
func (this *Engine) NewMotor(pinA int, pinB int, pinE int, reversed bool) *Motor {
	this.registerPin(pinA)
	this.registerPin(pinB)
	this.registerPin(pinE)
	return NewMotor(pinA, pinB, pinE, reversed)
}

// Returns a new instance of IRSensor.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
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
