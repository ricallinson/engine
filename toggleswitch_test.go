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

func TestToggleSwitch(t *testing.T) {

	var e *Engine

	BeforeEach(func() {
		e = Start(true)
	})

	AfterEach(func() {
		e.Stop()
	})

	Describe("NewToggleSwitch()", func() {
		It("should return an instance of ToggleSwitch", func() {
			AssertEqual(reflect.TypeOf(e.NewToggleSwitch(1)).String(), "*engine.ToggleSwitch")
		})
		It("should return have a pin mode of gpio.Input", func() {
			in := e.NewToggleSwitch(1)
			AssertEqual(in.GetMode(), gpio.Input)
		})
	})

	Report(t)
}
