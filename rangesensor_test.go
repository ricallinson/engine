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

func TestRangeSensor(t *testing.T) {

	var e *Engine

	BeforeEach(func() {
		e = Start(true)
	})

	AfterEach(func() {
		e.Stop()
	})

	Describe("NewRangeSensor()", func() {
		It("should return an instance of RangeSensor", func() {
			AssertEqual(reflect.TypeOf(NewRangeSensor(1, 2)).String(), "*engine.RangeSensor")
		})
		It("should return have a pin mode of rpio.Input", func() {
			in := NewRangeSensor(1, 2)
			trigger, echo := in.Pins()
			AssertEqual(rpio.StoredPinMode(rpio.Pin(trigger)), rpio.Output)
			AssertEqual(rpio.StoredPinMode(rpio.Pin(echo)), rpio.Input)
		})
	})

	Report(t)
}