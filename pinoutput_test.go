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

func TestPinOutput(t *testing.T) {

	var e *Engine

	BeforeEach(func() {
		e = Start(true)
	})

	AfterEach(func() {
		e.Stop()
	})

	Describe("NewPinOutput()", func() {
		It("should return an instance of PinOutput", func() {
			AssertEqual(reflect.TypeOf(NewPinOutput(1)).String(), "*engine.PinOutput")
		})
		It("should return have a pin mode of rpio.Input", func() {
			out := NewPinOutput(1)
			AssertEqual(rpio.StoredPinMode(rpio.Pin(out.PinNumber())), rpio.Output)
		})
	})

	Describe("Set()", func() {
		It("should return a value of rpio.High", func() {
			out := NewPinOutput(1)
			out.Set(1)
			AssertEqual(rpio.ReadPin(rpio.Pin(out.PinNumber())), rpio.High)
		})
		It("should return a value of rpio.Low", func() {
			out := NewPinOutput(1)
			out.Set(0)
			AssertEqual(rpio.ReadPin(rpio.Pin(out.PinNumber())), rpio.Low)
		})
		It("should return a value of 50%", func() {
			out := NewPinOutput(1)
			out.Set(0.5)
			AssertEqual(rpio.StoredPinPWM(rpio.Pin(out.PinNumber())), 50)
		})
		It("should return a value of rpio.Low from -0.1", func() {
			out := NewPinOutput(1)
			out.Set(-0.1)
			AssertEqual(rpio.ReadPin(rpio.Pin(out.PinNumber())), rpio.Low)
		})
	})

	Report(t)
}
