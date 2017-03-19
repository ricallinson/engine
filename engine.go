package engine

import (
	"log"
	"github.com/ricallinson/engine/rpio"
)

// Global list of used pins.
var PINS = make([]bool, 26, 26)

// Global engine in use flag.
var INUSE bool

type Engine struct {
	testing bool
}

func Start() *Engine {
	if INUSE {
		log.Panic("The engine is already being used.")
	}
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		// log.Panic("Could not access GPIO memory.", err)
	}
	this := &Engine{}
	INUSE = true
	return this
}

func (this *Engine) Stop() {
	PINS = make([]bool, 26, 26)
	INUSE = false
	rpio.Close()
}

func (this *Engine) NewLED(pin int) *LED {
	this.registerPin(pin)
	return NewLED(pin)
}

func (this *Engine) NewMotor(pin int, reversed bool) *Motor {
	this.registerPin(pin)
	return NewMotor(pin, reversed)
}

func (this *Engine) NewIRSensor(pin int) *IRSensor {
	this.registerPin(pin)
	return NewIRSensor(pin)
}

// If a pin has already been used then it results in a fatal error.
func (this *Engine) registerPin(pin int) {
	if PINS[pin] {
		log.Panic("Pin", pin, " has already been used.")
	}
	PINS[pin] = true
}
