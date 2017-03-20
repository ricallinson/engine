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

func TestMotor(t *testing.T) {

	var e *Engine

	BeforeEach(func() {
		e = Start(true)
	})

	AfterEach(func() {
		e.Stop()
	})

	Describe("NewMotor()", func() {
		It("should return an instance of Motor", func() {
			AssertEqual(reflect.TypeOf(NewMotor(1, 2, 3, false)).String(), "*engine.Motor")
		})
		It("should return have a pin mode of rpio.Output", func() {
			m := NewMotor(1, 2, 3, false)
			a, b, e := m.Pin()
			AssertEqual(rpio.MockGetPinMode(rpio.Pin(a)), rpio.Output)
			AssertEqual(rpio.MockGetPinMode(rpio.Pin(b)), rpio.Output)
			AssertEqual(rpio.MockGetPinMode(rpio.Pin(e)), rpio.Output)
		})
	})

	Describe("Set()", func() {
		It("should return a value of rpio.Low from zero", func() {
			m := NewMotor(1, 2, 3, false)
			m.Set(0)
			_, _, e := m.Pin()
			AssertEqual(rpio.ReadPin(rpio.Pin(e)), rpio.Low)
		})
		It("should return a value of rpio.High, rpio.Low, rpio.High from 0.1", func() {
			m := NewMotor(1, 2, 3, false)
			m.Set(0.1)
			a, b, e := m.Pin()
			AssertEqual(rpio.ReadPin(rpio.Pin(a)), rpio.High)
			AssertEqual(rpio.ReadPin(rpio.Pin(b)), rpio.Low)
			AssertEqual(rpio.ReadPin(rpio.Pin(e)), rpio.High)
		})
		It("should return a value of rpio.Low, rpio.High, rpio.High from -0.1", func() {
			m := NewMotor(1, 2, 3, false)
			m.Set(-0.1)
			a, b, e := m.Pin()
			AssertEqual(rpio.ReadPin(rpio.Pin(a)), rpio.Low)
			AssertEqual(rpio.ReadPin(rpio.Pin(b)), rpio.High)
			AssertEqual(rpio.ReadPin(rpio.Pin(e)), rpio.High)
		})
		It("should return a value of rpio.High, rpio.Low, rpio.High from -0.1", func() {
			m := NewMotor(1, 2, 3, true)
			m.Set(-0.1)
			a, b, e := m.Pin()
			AssertEqual(rpio.ReadPin(rpio.Pin(a)), rpio.High)
			AssertEqual(rpio.ReadPin(rpio.Pin(b)), rpio.Low)
			AssertEqual(rpio.ReadPin(rpio.Pin(e)), rpio.High)
		})
		It("should return a value of rpio.Low, rpio.High, rpio.High from 0.1", func() {
			m := NewMotor(1, 2, 3, true)
			m.Set(0.1)
			a, b, e := m.Pin()
			AssertEqual(rpio.ReadPin(rpio.Pin(a)), rpio.Low)
			AssertEqual(rpio.ReadPin(rpio.Pin(b)), rpio.High)
			AssertEqual(rpio.ReadPin(rpio.Pin(e)), rpio.High)
		})
	})

	Describe("and Stop(), Forwards() and Backwards()", func() {
		It("should return a value of rpio.Low from zero", func() {
			m := NewMotor(1, 2, 3, false)
			m.Stop()
			_, _, e := m.Pin()
			AssertEqual(rpio.ReadPin(rpio.Pin(e)), rpio.Low)
		})
		It("should return a value of rpio.High, rpio.Low, rpio.High from 0.1", func() {
			m := NewMotor(1, 2, 3, false)
			m.Forwards()
			a, b, e := m.Pin()
			AssertEqual(rpio.ReadPin(rpio.Pin(a)), rpio.High)
			AssertEqual(rpio.ReadPin(rpio.Pin(b)), rpio.Low)
			AssertEqual(rpio.ReadPin(rpio.Pin(e)), rpio.High)
		})
		It("should return a value of rpio.Low, rpio.High, rpio.High from -0.1", func() {
			m := NewMotor(1, 2, 3, false)
			m.Backwards()
			a, b, e := m.Pin()
			AssertEqual(rpio.ReadPin(rpio.Pin(a)), rpio.Low)
			AssertEqual(rpio.ReadPin(rpio.Pin(b)), rpio.High)
			AssertEqual(rpio.ReadPin(rpio.Pin(e)), rpio.High)
		})
	})

	Report(t)
}
