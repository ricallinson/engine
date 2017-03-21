//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package engine

import (
	"github.com/ricallinson/engine/rpio"
	"log"
	"time"
)

type RangeSensor struct {
	pinTrigger rpio.Pin
	pinEcho    rpio.Pin
}

// Returns a new instance of RangeSensor.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
// Controls a HC-SR04 Ultrasonic Range Sensor.
func NewRangeSensor(pinTrigger, pinEcho int) *RangeSensor {
	this := &RangeSensor{
		pinTrigger: rpio.Pin(pinTrigger),
		pinEcho:    rpio.Pin(pinEcho),
	}
	this.pinTrigger.Output()
	this.pinTrigger.Low()
	this.pinEcho.Input()
	// Wait for the sensor to settle.
	time.Sleep(2 * time.Second)
	return this
}

// Returns the pins that this instance is controlled by.
func (this *RangeSensor) Pins() (int, int) {
	return int(this.pinTrigger), int(this.pinEcho)
}

func (this *RangeSensor) Get() float32 {
	// The HC-SR04 sensor requires a short 10uS pulse to trigger the module, which will
	// cause the sensor to start the ranging program (8 ultrasound bursts at 40 kHz) in
	// order to obtain an echo response. So, to create our trigger pulse, we set out
	// trigger pin high for 10uS then set it low again.
	this.pinTrigger.High()
	time.Sleep(10 * time.Microsecond)
	this.pinTrigger.Low()

	// Our first step is to record the last rpio.Low timestamp for ECHO (pulse_start)
	// e.g. just before the return signal is received and the pin goes rpio.High.
	var pulseStart time.Time
	for this.pinEcho.Read() == rpio.Low {
		// Start the timer.
		pulseStart = time.Now()
	}

	// Once a signal is received, the value changes from low (0) to high (1), and the
	// signal will remain high for the duration of the echo pulse. We therefore also need
	// the last high timestamp for ECHO (pulse_end).
	var pulseDuration time.Duration
	for this.pinEcho.Read() == rpio.High {
		// Time taken for sound to travel to an obstacle (there and back divided by two).
		pulseDuration = time.Since(pulseStart)
	}

	// Distance in cm (at sea level air).
	distance := pulseDuration.Seconds() * 17150

	// Maybe round to two places?
	this.log(float32(distance))
	return float32(distance)
}

// Logs state of the assigned pin.
func (this *RangeSensor) log(cm float32) {
	log.Print("RangeSensor on pin ", this.pinTrigger, " and ", this.pinEcho, " measured a distance of ", cm, "cm")
}
