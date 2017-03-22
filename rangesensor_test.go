//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package engine

import (
	"fmt"
	"github.com/ricallinson/engine/rpio"
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
	"time"
)

func mockRangeSensor(pinTrigger int, pinEcho int, cm int) {
	trigger := rpio.Pin(pinTrigger)
	echo := rpio.Pin(pinEcho)
	// If the trigger is already rpio.High then fail.
	if trigger.Read() == rpio.High {
		fmt.Println("ERROR: Mock Range Sensor trigger was already on.")
		AssertEqual(true, false)
		return
	}
	// If the echo is already rpio.High then fail.
	if echo.Read() == rpio.High {
		fmt.Println("ERROR: Mock Range Sensor echo was already on.")
		AssertEqual(true, false)
		return
	}
	// The sound over distance measurement at sea level is 3430cm per 1 second.
	timeInMicroseconds := (1000 * 1000 / 3430) * cm
	// Listen for the tigger signal.
	for trigger.Read() != rpio.High {
		// No code here.
	}
	// Loop 8 times to mimic the HC-SR04 Ultrasonic Range Sensors operation.
	fmt.Println("Mock Range Sensor was activated.")
	for i := 0; i < 8; i++ {
		// Create the echo signal by waiting for timeInMicroseconds.
		echo.Write(rpio.High)
		time.Sleep(time.Duration(timeInMicroseconds) * time.Microsecond)
		echo.Write(rpio.Low)
		time.Sleep(time.Millisecond)
	}
}

func TestRangeSensor(t *testing.T) {

	var e *Engine

	BeforeEach(func() {
		e = Start(true)
	})

	AfterEach(func() {
		e.Stop()
	})

	Describe("NewRangeSensor()", func() {
		It("should return an instance of RangeSensor", func() {
			AssertEqual(reflect.TypeOf(NewRangeSensor(1, 2)).String(), "*engine.RangeSensor")
		})
		It("should return have a pin mode of rpio.Input", func() {
			rs := NewRangeSensor(1, 2)
			trigger, echo := rs.Pins()
			AssertEqual(rpio.StoredPinMode(rpio.Pin(trigger)), rpio.Output)
			AssertEqual(rpio.StoredPinMode(rpio.Pin(echo)), rpio.Input)
		})
	})

	Describe("Get()", func() {
		It("should return a distance of 0cm for 100cm", func() {
			rs := NewRangeSensor(1, 2)
			go mockRangeSensor(1, 2, 100)
			time.Sleep(100 * time.Microsecond)
			d := rs.Get()
			AssertEqual(d, float32(0))
		})

		It("should return a distance of 1cm", func() {
			rs := NewRangeSensor(1, 2)
			go mockRangeSensor(1, 2, 1)
			time.Sleep(100 * time.Microsecond)
			d := rs.Get()
			if d > 1 {
				AssertEqual(d, d)
			} else {
				AssertEqual(d, float32(1))
			}
		})

		It("should return a distance of 15cm", func() {
			rs := NewRangeSensor(1, 2)
			go mockRangeSensor(1, 2, 15)
			time.Sleep(100 * time.Microsecond)
			d := rs.Get()
			if d > 15 {
				AssertEqual(d, d)
			} else {
				AssertEqual(d, float32(15))
			}
		})

		It("should return a distance of 30cm", func() {
			rs := NewRangeSensor(1, 2)
			go mockRangeSensor(1, 2, 30)
			time.Sleep(100 * time.Microsecond)
			d := rs.Get()
			if d > 30 {
				AssertEqual(d, d)
			} else {
				AssertEqual(d, float32(30))
			}
		})
	})

	Report(t)
}
