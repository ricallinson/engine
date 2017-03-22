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
	time.Sleep(1 * time.Second)
	return this
}

// Returns the pins that this instance is controlled by.
func (this *RangeSensor) Pins() (int, int) {
	return int(this.pinTrigger), int(this.pinEcho)
}

func (this *RangeSensor) Get() float32 {
	// Create an array to hold the measurement values.
	// This is done before trigging the sensor as timing is critical to a good measurement.
	distances := make([]time.Duration, 8, 8)
	// The HC-SR04 sensor requires a short 10uS pulse to trigger the module, which will
	// cause the sensor to start the ranging program (8 ultrasound bursts at 40 kHz) in
	// order to obtain an echo response. So, to create our trigger pulse, we set out
	// trigger pin high for 10uS then set it low again.
	this.pinTrigger.High()
	time.Sleep(11 * time.Microsecond)
	this.pinTrigger.Low()
	// Measure the distance 8 times.
	for i := 7; i >= 0; i-- {
		// distances[i] = this.takeMeasurement()

		// TEST START
		// Our first step is to record the last rpio.Low timestamp for pinEcho (pulseStart)
		// e.g. just before the return signal is received and pinEcho goes rpio.High.
		var pulseStart time.Time
		for this.pinEcho.Read() == rpio.Low {
			// Start the timer.
			pulseStart = time.Now()
		}
		// Once a signal is received, the value changes from rpio.Low (0) to rpio.High (1), and the
		// signal will remain rpio.High for the duration of the pinEcho pulse. We therefore also need
		// the last rpio.High timestamp for pinEcho to give us a duration.
		// var pulseDuration time.Duration
		for this.pinEcho.Read() == rpio.High {
			// Time taken for sound to travel to an obstacle (there and back divided by two).
			distances[i] = time.Since(pulseStart)
		}
		// TEST END
	}
	// Only use measurements within range.
	valid := 0
	distance := float32(0)
	for _, pulseDuration := range distances {
		// The sound over distance measurement at sea level is 3430cm per 1 second.
		measurement := float32(pulseDuration.Seconds() * 3430)
		if measurement > 0.1 && measurement < 31 {
			distance += measurement
			valid++
		}
	}
	// If there were no valid measurements then return 0.
	if valid == 0 {
		return 0
	}
	// Take an average of the in range measurements.
	distance = distance / float32(valid)
	// Log the final measurement and return it.
	this.log(distance)
	return distance
}

func (this *RangeSensor) takeMeasurement() time.Duration {
	// Our first step is to record the last rpio.Low timestamp for pinEcho (pulseStart)
	// e.g. just before the return signal is received and pinEcho goes rpio.High.
	var pulseStart time.Time
	for this.pinEcho.Read() == rpio.Low {
		// Start the timer.
		pulseStart = time.Now()
	}
	// Once a signal is received, the value changes from rpio.Low (0) to rpio.High (1), and the
	// signal will remain rpio.High for the duration of the pinEcho pulse. We therefore also need
	// the last rpio.High timestamp for pinEcho to give us a duration.
	var pulseDuration time.Duration
	for this.pinEcho.Read() == rpio.High {
		// Time taken for sound to travel to an obstacle (there and back divided by two).
		pulseDuration = time.Since(pulseStart)
	}
	return pulseDuration
}

// Logs state of the assigned pin.
func (this *RangeSensor) log(cm float32) {
	log.Print("RangeSensor on pin ", this.pinTrigger, " and ", this.pinEcho, " measured a distance of ", cm, "cm.")
}
