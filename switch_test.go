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

func TestSwitch(t *testing.T) {

	var e *Engine

	BeforeEach(func() {
		e = Start(true)
	})

	AfterEach(func() {
		e.Stop()
	})

	Describe("NewSwitch()", func() {
		It("should return an instance of Switch", func() {
			AssertEqual(reflect.TypeOf(NewSwitch(1)).String(), "*engine.Switch")
		})
		It("should return have a pin mode of rpio.Input", func() {
			in := NewSwitch(1)
			AssertEqual(rpio.StoredPinMode(rpio.Pin(in.PinNumber())), rpio.Input)
		})
	})

	Report(t)
}
