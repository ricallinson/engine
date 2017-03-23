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

func mockRangeSensor(pinTrigger int, pinEcho int, cm float32) {
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
	// Listen for the trigger signal.
	for trigger.Read() != rpio.High {
		// No code here.
	}
	fmt.Println("Mock Range Sensor was activated.")
	// Mock the sensors 40hz signal sent 8 times. Formual is 40hz = 25000uS * 8.
	time.Sleep(25000 * 8 * time.Microsecond)
	// The formula for distance measured is cm = uS / 58.
	timeInMicroseconds := time.Duration(cm * 580)
	// Create the echo signal by waiting for timeInMicroseconds.
	echo.Write(rpio.High)
	time.Sleep(timeInMicroseconds * time.Microsecond)
	echo.Write(rpio.Low)
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
		// It("should return a distance of 1cm for 0cm", func() {
		// 	cm := float32(1)
		// 	rs := NewRangeSensor(1, 2)
		// 	go mockRangeSensor(1, 2, cm)
		// 	time.Sleep(100 * time.Microsecond)
		// 	d := rs.Get()
		// 	AssertEqual(d, float32(1))
		// })

		It("should return a distance of 0cm for 500cm", func() {
			cm := float32(500)
			rs := NewRangeSensor(1, 2)
			go mockRangeSensor(1, 2, cm)
			time.Sleep(100 * time.Microsecond)
			d := rs.Get()
			AssertEqual(d, float32(400))
		})

		It("should return a distance of 10cm +/- 20%", func() {
			cm := float32(10)
			rs := NewRangeSensor(1, 2)
			go mockRangeSensor(1, 2, cm)
			time.Sleep(100 * time.Microsecond)
			d := rs.Get()
			if d > cm*0.8 && d < cm*1.2 {
				AssertEqual(d, d)
			} else {
				AssertEqual(d, float32(cm))
			}
		})

		It("should return a distance of 50cm +/- 20%", func() {
			cm := float32(50)
			rs := NewRangeSensor(1, 2)
			go mockRangeSensor(1, 2, cm)
			time.Sleep(100 * time.Microsecond)
			d := rs.Get()
			if d > cm*0.8 && d < cm*1.2 {
				AssertEqual(d, d)
			} else {
				AssertEqual(d, float32(cm))
			}
		})

		It("should return a distance of 100cm +/- 20%", func() {
			cm := float32(100)
			rs := NewRangeSensor(1, 2)
			go mockRangeSensor(1, 2, cm)
			time.Sleep(100 * time.Microsecond)
			d := rs.Get()
			if d > cm*0.8 && d < cm*1.2 {
				AssertEqual(d, d)
			} else {
				AssertEqual(d, float32(cm))
			}
		})
	})

	Report(t)
}
