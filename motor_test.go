//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD Licensengine.
// See the accompanying LICENSE file for terms.
//

package engine

import (
	"github.com/ricallinson/engine/gpio"
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
)

func TestMotor(t *testing.T) {

	var engine *Engine

	BeforeEach(func() {
		engine = Start(true)
	})

	AfterEach(func() {
		engine.Stop()
	})

	Describe("engine.NewMotor()", func() {
		It("should return an instance of Motor", func() {
			AssertEqual(reflect.TypeOf(engine.NewMotor(1, 2, 3, false)).String(), "*engine.Motor")
		})
		It("should return have a pin mode of gpio.Output", func() {
			m := engine.NewMotor(1, 2, 3, false)
			a, b, e := m.PinsOut()
			AssertEqual(engine.GetGpioPin(a).GetMode(), gpio.Output)
			AssertEqual(engine.GetGpioPin(b).GetMode(), gpio.Output)
			AssertEqual(engine.GetGpioPin(e).GetMode(), gpio.Output)
		})
	})

	Describe("Set()", func() {
		It("should return a value of gpio.Low from zero", func() {
			m := engine.NewMotor(1, 2, 3, false)
			m.Set(0)
			_, _, e := m.PinsOut()
			AssertEqual(engine.GetGpioPin(e).LastWrite(), gpio.Low)
		})
		It("should return a value of gpio.High, gpio.Low, 10% from 0.1", func() {
			m := engine.NewMotor(1, 2, 3, false)
			m.Set(0.1)
			a, b, e := m.PinsOut()
			AssertEqual(engine.GetGpioPin(a).GetModulation(), 10)
			AssertEqual(engine.GetGpioPin(b).LastWrite(), gpio.Low)
			AssertEqual(engine.GetGpioPin(e).LastWrite(), gpio.High)
		})
		It("should return a value of gpio.Low, gpio.High, gpio.High from -0.1", func() {
			m := engine.NewMotor(1, 2, 3, false)
			m.Set(-0.1)
			a, b, e := m.PinsOut()
			AssertEqual(engine.GetGpioPin(a).LastWrite(), gpio.Low)
			AssertEqual(engine.GetGpioPin(b).GetModulation(), 10)
			AssertEqual(engine.GetGpioPin(e).LastWrite(), gpio.High)
		})
		It("should return a value of gpio.High, gpio.Low, 10% from -0.1", func() {
			m := engine.NewMotor(1, 2, 3, true)
			m.Set(-0.1)
			a, b, e := m.PinsOut()
			AssertEqual(engine.GetGpioPin(a).GetModulation(), 10)
			AssertEqual(engine.GetGpioPin(b).LastWrite(), gpio.Low)
			AssertEqual(engine.GetGpioPin(e).LastWrite(), gpio.High)
		})
		It("should return a value of gpio.Low, gpio.High, 10% from 0.1", func() {
			m := engine.NewMotor(1, 2, 3, true)
			m.Set(0.1)
			a, b, e := m.PinsOut()
			AssertEqual(engine.GetGpioPin(a).LastWrite(), gpio.Low)
			AssertEqual(engine.GetGpioPin(b).GetModulation(), 10)
			AssertEqual(engine.GetGpioPin(e).LastWrite(), gpio.High)
		})
	})

	Describe("and Stop(), Forwards() and Backwards()", func() {
		It("should return a value of gpio.Low from Stop()", func() {
			m := engine.NewMotor(1, 2, 3, false)
			m.Stop()
			_, _, e := m.PinsOut()
			AssertEqual(engine.GetGpioPin(e).LastWrite(), gpio.Low)
		})
		It("should return a value of gpio.High, gpio.Low, gpio.High from Forwards()", func() {
			m := engine.NewMotor(1, 2, 3, false)
			m.Forwards()
			a, b, e := m.PinsOut()
			AssertEqual(engine.GetGpioPin(a).LastWrite(), gpio.High)
			AssertEqual(engine.GetGpioPin(b).LastWrite(), gpio.Low)
			AssertEqual(engine.GetGpioPin(e).LastWrite(), gpio.High)
		})
		It("should return a value of gpio.Low, gpio.High, gpio.High from Backwards()", func() {
			m := engine.NewMotor(1, 2, 3, false)
			m.Backwards()
			a, b, e := m.PinsOut()
			AssertEqual(engine.GetGpioPin(a).LastWrite(), gpio.Low)
			AssertEqual(engine.GetGpioPin(b).LastWrite(), gpio.High)
			AssertEqual(engine.GetGpioPin(e).LastWrite(), gpio.High)
		})
	})

	Report(t)
}
