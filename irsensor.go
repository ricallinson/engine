package engine

import (
	"github.com/ricallinson/engine/rpio"
	"log"
)

type IRSensor struct {
	pin rpio.Pin
}

func NewIRSensor(pin int) *IRSensor {
	this := &IRSensor{
		pin: rpio.Pin(pin),
	}
	this.pin.Input()
	return this
}

func (this *IRSensor) Get() float32 {
	val := this.pin.Read()
	this.log(val)
	return float32(val)
}

func (this *IRSensor) log(val rpio.State) {
	log.Print("IRSensor on pin ", this.pin, " read a value of ", val)
}
