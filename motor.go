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

type Motor struct {
	pinA      rpio.Pin
	pinB      rpio.Pin
	pinE      rpio.Pin
	direction int
}

// Returns a new instance of Motor.
// The value of `pinX` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
// If `reversed` is `true` then forwards and backwards will be flipped.
// Controls a L293D Stepper Motor Driver chip.
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

// Returns the pins that this instance is controlled by.
func (this *Motor) PinNumber() (int, int, int) {
	return int(this.pinA), int(this.pinB), int(this.pinE)
}

// Sets the motor full power forward.
func (this *Motor) Forwards() {
	this.Set(1)
}

// Sets the motor full power backwards.
func (this *Motor) Backwards() {
	this.Set(-1)
}

// Removes all power to motor.
func (this *Motor) Stop() {
	this.Set(0)
}

// Set the current value of this instances Motor.
// The range is -1 to 1 rounded up where -1 is full power backwards, 0 is stop and 1 is full power forwards.
func (this *Motor) Set(val float32) {
	val = val * float32(this.direction)
	if val == 0 {
		// Stop.
		this.pinA.Low()
		this.pinB.Low()
		this.pinE.Low()
	} else if val > 0 {
		// Forwards.
		this.pinA.WritePWM(int(val * 100))
		this.pinB.Low()
		this.pinE.High()
	} else {
		// Backwards.
		this.pinA.Low()
		this.pinB.WritePWM(int(val * -1 * 100))
		this.pinE.High()
	}
	this.log()
}

// Logs state of the assigned pins.
func (this *Motor) log() {
	log.Print("Motor on pins ", this.pinA, ", ", this.pinB, " and ", this.pinE, " set values to ", this.pinA.Read(), ", ", this.pinB.Read(), " and ", this.pinE.Read())
}
