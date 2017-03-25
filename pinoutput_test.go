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

func TestPinOutput(t *testing.T) {

	var e *Engine

	BeforeEach(func() {
		e = Start(true)
	})

	AfterEach(func() {
		e.Stop()
	})

	Describe("e.NewPinOutput()", func() {
		It("should return an instance of PinOutput", func() {
			AssertEqual(reflect.TypeOf(e.NewPinOutput(1)).String(), "*engine.PinOutput")
		})
		It("should return a pin mode of gpio.Input", func() {
			out := e.NewPinOutput(1)
			AssertEqual(out.GetMode(), gpio.Output)
		})
	})

	Describe("Set()", func() {
		It("should return a value of gpio.High", func() {
			out := e.NewPinOutput(1)
			out.Set(1)
			AssertEqual(out.LastWrite(), gpio.High)
		})
		It("should return a value of gpio.Low", func() {
			out := e.NewPinOutput(1)
			out.Set(0)
			AssertEqual(out.LastWrite(), gpio.Low)
		})
		It("should return a value of 50%", func() {
			out := e.NewPinOutput(1)
			out.Set(0.5)
			AssertEqual(out.GetModulation(), 50)
		})
		It("should return a value of gpio.Low from -0.1", func() {
			out := e.NewPinOutput(1)
			out.Set(-0.1)
			AssertEqual(out.LastWrite(), gpio.Low)
		})
	})

	Report(t)
}
