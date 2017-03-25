//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package engine

import (
	"github.com/ricallinson/engine/gpio"
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
)

func TestLED(t *testing.T) {

	var e *Engine

	BeforeEach(func() {
		e = Start(true)
	})

	AfterEach(func() {
		e.Stop()
	})

	Describe("NewLED()", func() {
		It("should return an instance of LED", func() {
			AssertEqual(reflect.TypeOf(e.NewLED(1)).String(), "*engine.LED")
		})
		It("should return have a pin mode of rpio.Output", func() {
			led := e.NewLED(1)
			AssertEqual(led.GetMode(), gpio.Output)
		})
	})

	Describe("On() Off()", func() {
		It("should return a value of rpio.High and then rpio.Low", func() {
			led := e.NewLED(1)
			led.On()
			AssertEqual(led.LastWrite(), gpio.High)
			led.Off()
			AssertEqual(led.LastWrite(), gpio.Low)
		})
	})

	Describe("Toggle()", func() {
		It("should return a value of rpio.High and then rpio.Low", func() {
			led := e.NewLED(1)
			led.Toggle()
			AssertEqual(led.LastWrite(), gpio.High)
			led.Toggle()
			AssertEqual(led.LastWrite(), gpio.Low)
		})
	})

	Report(t)
}
