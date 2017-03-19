package engine

import (
	"github.com/ricallinson/engine/rpio"
	"log"
)

type Motor struct {
	pinOne    rpio.Pin
	pinTwo    rpio.Pin
	direction int
}

func NewMotor(pinOne int, pinTwo int, reversed bool) *Motor {
	this := &Motor{
		pinOne:    rpio.Pin(pinOne),
		pinTwo:    rpio.Pin(pinTwo),
		direction: 1,
	}
	if reversed {
		this.direction = -1
	}
	this.pinOne.Output()
	this.pinTwo.Output()
	return this
}

func (this *Motor) Forwards() {
	this.Set(1)
}

func (this *Motor) Backwards() {
	this.Set(-1)
}

func (this *Motor) Stop() {
	this.Set(0)
}

func (this *Motor) Set(val float32) {
	val = val * float32(this.direction)
	if val == 0 {
		this.pinOne.Low()
		this.pinTwo.Low()
	} else if val > 0 {
		this.pinOne.High()
		this.pinTwo.Low()
	} else {
		this.pinOne.Low()
		this.pinTwo.High()
	}
	this.log()
}

func (this *Motor) log() {
	log.Print("Motor on pins ", this.pinOne, " and ", this.pinTwo, " set values to ", this.pinOne.Read(), " and ", this.pinTwo.Read())
}
