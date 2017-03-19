package engine

import (
	"log"
)

// Global list of used pins.
var PINS = make([]bool, 26, 26)
var INUSE bool

type Engine struct {
}

func Start() *Engine {
	if INUSE {
		log.Panic("The engine is already being used.")
	}
	this := &Engine{}
	INUSE = true
	return this
}

func (this *Engine) Stop() {
	PINS = make([]bool, 26, 26)
	INUSE = false
}

func (this *Engine) NewLED(pin int) *LED {
	this.registerPin(pin)
	led := &LED{}
	return led
}

func (this *Engine) NewMotor(pin int, reversed bool) *Motor {
	this.registerPin(pin)
	motor := &Motor{}
	return motor
}

func (this *Engine) NewIRSensor(pin int) *IRSensor {
	this.registerPin(pin)
	ir := &IRSensor{}
	return ir
}

// If a pin has already been used then it results in a fatal error.
func (this *Engine) registerPin(pin int) {
	if PINS[pin] {
		log.Panic("Pin", pin, " has already been used.")
	}
	PINS[pin] = true
}
