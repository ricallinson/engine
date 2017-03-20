package engine

import (
	"github.com/ricallinson/engine/rpio"
	"log"
)

type Motor struct {
	pinA      rpio.Pin
	pinB      rpio.Pin
	pinE      rpio.Pin
	direction int
}

func NewMotor(pinA int, pinB int, pinE int, reversed bool) *Motor {
	this := &Motor{
		pinA:      rpio.Pin(pinA),
		pinB:      rpio.Pin(pinB),
		pinE:      rpio.Pin(pinE),
		direction: 1,
	}
	if reversed {
		this.direction = -1
	}
	this.pinA.Output()
	this.pinB.Output()
	this.pinE.Output()
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
		this.pinE.Low()
	} else if val > 0 {
		this.pinA.High()
		this.pinB.Low()
		this.pinE.High()
	} else {
		this.pinA.Low()
		this.pinB.High()
		this.pinE.High()
	}
	this.log()
}

func (this *Motor) log() {
	log.Print("Motor on pins ", this.pinA, ", ", this.pinB, " and ", this.pinE, " set values to ", this.pinA.Read(), ", ", this.pinB.Read(), " and ", this.pinE.Read())
}
