package engine

import (
	"github.com/ricallinson/engine/rpio"
)

type LED struct {
	Pin rpio.Pin
}

func NewLED(pin int) *LED {
	this := &LED{
		Pin: rpio.Pin(pin),
	}
	this.Pin.Output()
	return this
}

func (this *LED) Toggle() {
	this.Pin.Toggle()
}

func (this *LED) Set(val float32) {

}
