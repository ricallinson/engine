//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package engine

import (
	"github.com/ricallinson/engine/gpio"
	"log"
)

// Interface for managing hardware pins and instantiating control objects such as Motors, LEDs and IRSensors.
type Engine struct {
	gpio     *gpio.GpioSingleton
	pinsUsed []bool
}

// Starts the engine returning an instance of Engine.
// Subsequent calls to Start() without calling Stop() first will generate a panic.
func Start(mock bool) *Engine {
	this := &Engine{
		gpio:     gpio.Gpio(),
		pinsUsed: make([]bool, 26, 26),
	}
	return this
}

// Stops the engine and releases all resources.
// It must be called before Start() can be called again.
func (this *Engine) Stop() {
	this.pinsUsed = make([]bool, 26, 26)
	this.gpio.Close()
}

// Returns the GpioPin for the given pin.
func (this *Engine) GetGpioPin(pin uint8) *gpio.GpioPin {
	return this.gpio.Pin(pin)
}

// Returns a new instance of a Pin.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
func (this *Engine) NewPin(pin uint8) *Pin {
	this.registerPin(pin)
	p := this.gpio.Pin(pin)
	return NewPin(p)
}

// Returns a new instance of a PinInput.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
func (this *Engine) NewPinInput(pin uint8) *PinInput {
	this.registerPin(pin)
	p := this.gpio.Pin(pin)
	return NewPinInput(p)
}

// Returns a new instance of a PinOutput.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
func (this *Engine) NewPinOutput(pin uint8) *PinOutput {
	this.registerPin(pin)
	p := this.gpio.Pin(pin)
	return NewPinOutput(p)
}

// Returns a new instance of a ToggleSwitch.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
func (this *Engine) NewToggleSwitch(pin uint8) *ToggleSwitch {
	this.registerPin(pin)
	p := this.gpio.Pin(pin)
	return NewToggleSwitch(p)
}

// Returns a new instance of a LED.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
func (this *Engine) NewLED(pin uint8) *LED {
	this.registerPin(pin)
	p := this.gpio.Pin(pin)
	return NewLED(p)
}

// Returns a new instance of an IRSensor.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
func (this *Engine) NewIRSensor(pin uint8) *IRSensor {
	this.registerPin(pin)
	p := this.gpio.Pin(pin)
	return NewIRSensor(p)
}

// Returns a new instance of a RangeSensor.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
func (this *Engine) NewRangeSensor(pinTrigger uint8, pinEcho uint8) *RangeSensor {
	this.registerPin(pinTrigger)
	this.registerPin(pinEcho)
	t := this.gpio.Pin(pinTrigger)
	e := this.gpio.Pin(pinEcho)
	return NewRangeSensor(t, e)
}

// Returns a new instance of a Motor.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
func (this *Engine) NewMotor(pinA uint8, pinB uint8, pinE uint8, reversed bool) *Motor {
	this.registerPin(pinA)
	this.registerPin(pinB)
	this.registerPin(pinE)
	a := this.gpio.Pin(pinA)
	b := this.gpio.Pin(pinB)
	e := this.gpio.Pin(pinE)
	return NewMotor(a, b, e, reversed)
}

// If a pin has already been used then it results in a fatal error.
func (this *Engine) registerPin(pin uint8) {
	if pin >= uint8(len(this.pinsUsed)) || pin < 1 {
		log.Panic("Pin number ", pin, " is out of range.")
	}
	if this.pinsUsed[pin] {
		log.Panic("Pin number ", pin, " has already been used.")
	}
	this.pinsUsed[pin] = true
}
