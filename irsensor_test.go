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

func TestIRSensor(t *testing.T) {

	var e *Engine

	BeforeEach(func() {
		e = Start(true)
	})

	AfterEach(func() {
		e.Stop()
	})

	Describe("NewIRSensor()", func() {
		It("should return an instance of IRSensor", func() {
			AssertEqual(reflect.TypeOf(e.NewIRSensor(1)).String(), "*engine.IRSensor")
		})
		It("should return a pin mode of rpio.Input", func() {
			ir := e.NewIRSensor(1)
			AssertEqual(ir.GetMode(), gpio.Input)
		})
	})

	Describe("Get()", func() {
		It("should return a value of zero", func() {
			ir := e.NewIRSensor(1)
			ir.Low()
			AssertEqual(ir.Get(), float32(0))
		})
		It("should return a value of one", func() {
			ir := e.NewIRSensor(1)
			ir.High()
			AssertEqual(ir.Get(), float32(1))
		})
	})

	Report(t)
}
