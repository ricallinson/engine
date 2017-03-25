//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package engine

import (
	"fmt"
	"github.com/ricallinson/engine/gpio"
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
	"time"
)

func mockRangeSensor(trigger *gpio.GpioPin, echo *gpio.GpioPin, cm float32, failTrigger ...bool) {
	if len(failTrigger) > 0 {
		time.Sleep(1100 * time.Millisecond)
		return
	}
	// If the trigger is already gpio.High then fail.
	if trigger.Read() == gpio.High {
		fmt.Println("ERROR: Mock Range Sensor trigger was already on.")
		AssertEqual(true, false)
		return
	}
	// If the echo is already gpio.High then fail.
	if echo.Read() == gpio.High {
		fmt.Println("ERROR: Mock Range Sensor echo was already on.")
		AssertEqual(true, false)
		return
	}
	// Listen for the trigger signal.
	for trigger.Read() != gpio.High {
		// No code here.
	}
	fmt.Println("Mock Range Sensor was activated.")
	// Mock the sensors 40hz signal sent 8 times. Formual is (40hz == 25000uS) * 8.
	time.Sleep(25000 * 8 * time.Microsecond)
	// The formula for distance measured on the HC-SR04 sensor is cm = uS / 58.
	timeInMicroseconds := time.Duration(cm * 58)
	// Create the echo signal by waiting for timeInMicroseconds.
	echo.Write(gpio.High)
	time.Sleep(timeInMicroseconds * time.Microsecond)
	echo.Write(gpio.Low)
}

func TestRangeSensor(t *testing.T) {

	var e *Engine
	var lower = float32(0.5)
	var upper = float32(1.5)

	BeforeEach(func() {
		e = Start(true)
	})

	AfterEach(func() {
		e.Stop()
	})

	Describe("e.NewRangeSensor()", func() {
		It("should return an instance of RangeSensor", func() {
			AssertEqual(reflect.TypeOf(e.NewRangeSensor(1, 2)).String(), "*engine.RangeSensor")
		})
		It("should return a pin mode of gpio.Input", func() {
			rs := e.NewRangeSensor(1, 2)
			trigger, echo := rs.Pins()
			AssertEqual(e.GetGpioPin(trigger).GetMode(), gpio.Output)
			AssertEqual(e.GetGpioPin(echo).GetMode(), gpio.Input)
		})
	})

	Describe("Get()", func() {

		It("should return a distance of -1cm for 0cm", func() {
			cm := float32(0)
			rs := e.NewRangeSensor(1, 2)
			go mockRangeSensor(e.GetGpioPin(1), e.GetGpioPin(2), cm)
			time.Sleep(100 * time.Microsecond)
			d := rs.Get()
			AssertEqual(d, float32(-1))
		})

		It("should return a distance of 400cm for 500cm", func() {
			cm := float32(500)
			rs := e.NewRangeSensor(1, 2)
			go mockRangeSensor(e.GetGpioPin(1), e.GetGpioPin(2), cm)
			time.Sleep(100 * time.Microsecond)
			d := rs.Get()
			AssertEqual(d, float32(400))
		})

		It("should return a distance of 10cm +/-", func() {
			cm := float32(10)
			rs := e.NewRangeSensor(1, 2)
			go mockRangeSensor(e.GetGpioPin(1), e.GetGpioPin(2), cm)
			time.Sleep(100 * time.Microsecond)
			d := rs.Get()
			if d > cm*lower && d < cm*upper {
				AssertEqual(d, d)
			} else {
				AssertEqual(d, float32(cm))
			}
		})

		It("should return a distance of 50cm +/-", func() {
			cm := float32(50)
			rs := e.NewRangeSensor(1, 2)
			go mockRangeSensor(e.GetGpioPin(1), e.GetGpioPin(2), cm)
			time.Sleep(100 * time.Microsecond)
			d := rs.Get()
			if d > cm*lower && d < cm*upper {
				AssertEqual(d, d)
			} else {
				AssertEqual(d, float32(cm))
			}
		})

		It("should return a distance of 100cm +/-", func() {
			cm := float32(100)
			rs := e.NewRangeSensor(1, 2)
			go mockRangeSensor(e.GetGpioPin(1), e.GetGpioPin(2), cm)
			time.Sleep(100 * time.Microsecond)
			d := rs.Get()
			if d > cm*lower && d < cm*upper {
				AssertEqual(d, d)
			} else {
				AssertEqual(d, float32(cm))
			}
		})

		It("should return a distance of 390cm +/-", func() {
			cm := float32(390)
			rs := e.NewRangeSensor(1, 2)
			go mockRangeSensor(e.GetGpioPin(1), e.GetGpioPin(2), cm)
			time.Sleep(100 * time.Microsecond)
			d := rs.Get()
			if d > cm*lower && d < cm*upper {
				AssertEqual(d, d)
			} else {
				AssertEqual(d, float32(cm))
			}
		})

		It("should return a distance of 0cm as ranging never started", func() {
			cm := float32(10)
			rs := e.NewRangeSensor(1, 2)
			go mockRangeSensor(e.GetGpioPin(1), e.GetGpioPin(2), cm, true)
			time.Sleep(100 * time.Microsecond)
			d := rs.Get()
			AssertEqual(d, float32(-1))
		})

		It("should return a distance of 0cm as ranging took too long", func() {
			cm := float32(58000)
			rs := e.NewRangeSensor(1, 2)
			go mockRangeSensor(e.GetGpioPin(1), e.GetGpioPin(2), cm)
			time.Sleep(100 * time.Microsecond)
			d := rs.Get()
			AssertEqual(d, float32(-1))
		})
	})

	Report(t)
}
