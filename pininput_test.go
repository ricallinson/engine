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

func TestPinInput(t *testing.T) {

	var e *Engine

	BeforeEach(func() {
		e = Start(true)
	})

	AfterEach(func() {
		e.Stop()
	})

	Describe("e.NewPinInput()", func() {
		It("should return an instance of PinIntput", func() {
			AssertEqual(reflect.TypeOf(e.NewPinInput(1)).String(), "*engine.PinInput")
		})
		It("should return a pin mode of gpio.Input", func() {
			in := e.NewPinInput(1)
			in.Input()
			AssertEqual(in.GetMode(), gpio.Input)
		})
	})

	Describe("Set()", func() {
		It("should return a value of zero", func() {
			in := e.NewPinInput(1)
			AssertEqual(in.Get(), float32(0))
		})
		It("should return a value of one", func() {
			in := e.NewPinInput(1)
			in.Write(gpio.High)
			AssertEqual(in.Get(), float32(1))
		})
	})

	Report(t)
}
