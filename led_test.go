//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package engine

import (
	"github.com/ricallinson/engine/rpio"
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
			AssertEqual(reflect.TypeOf(NewLED(1)).String(), "*engine.LED")
		})
		It("should return have a pin mode of rpio.Output", func() {
			led := NewLED(1)
			AssertEqual(rpio.MockGetPinMode(rpio.Pin(led.Pin())), rpio.Output)
		})
	})

	Describe("Set()", func() {
		It("should return a value of rpio.High", func() {
			led := NewLED(1)
			led.Set(1)
			AssertEqual(rpio.ReadPin(rpio.Pin(led.Pin())), rpio.High)
		})
		It("should return a value of rpio.Low", func() {
			led := NewLED(1)
			led.Set(0)
			AssertEqual(rpio.ReadPin(rpio.Pin(led.Pin())), rpio.Low)
		})
		It("should return a value of rpio.High from 0.1", func() {
			led := NewLED(1)
			led.Set(0.5)
			AssertEqual(rpio.ReadPin(rpio.Pin(led.Pin())), rpio.High)
		})
		It("should return a value of rpio.Low from -0.1", func() {
			led := NewLED(1)
			led.Set(0)
			AssertEqual(rpio.ReadPin(rpio.Pin(led.Pin())), rpio.Low)
		})
	})

	Describe("Toggle()", func() {
		It("should return a value of rpio.High and then rpio.Low", func() {
			led := NewLED(1)
			led.Toggle()
			AssertEqual(rpio.ReadPin(rpio.Pin(led.Pin())), rpio.High)
			led.Toggle()
			AssertEqual(rpio.ReadPin(rpio.Pin(led.Pin())), rpio.Low)
		})
	})

	Report(t)
}
