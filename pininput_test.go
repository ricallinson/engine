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

func TestPinInput(t *testing.T) {

	var e *Engine

	BeforeEach(func() {
		e = Start(true)
	})

	AfterEach(func() {
		e.Stop()
	})

	Describe("NewPinInput()", func() {
		It("should return an instance of PinIntput", func() {
			AssertEqual(reflect.TypeOf(NewPinInput(1)).String(), "*engine.PinInput")
		})
		It("should return have a pin mode of rpio.Input", func() {
			in := NewPinInput(1)
			AssertEqual(rpio.StoredPinMode(rpio.Pin(in.PinNumber())), rpio.Input)
		})
	})

	Describe("Set()", func() {
		It("should return a value of zero", func() {
			in := NewPinInput(1)
			AssertEqual(in.Get(), float32(0))
		})
		It("should return a value of one", func() {
			in := NewPinInput(1)
			rpio.WritePin(rpio.Pin(in.PinNumber()), rpio.High)
			AssertEqual(in.Get(), float32(1))
		})
	})

	Report(t)
}
