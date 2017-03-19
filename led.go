package engine

import (
	"github.com/ricallinson/engine/rpio"
	"log"
)

type LED struct {
	pin rpio.Pin
}

func NewLED(pin int) *LED {
	this := &LED{
		pin: rpio.Pin(pin),
	}
	this.pin.Output()
	return this
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
