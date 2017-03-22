//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package engine

import (
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
)

func TestPin(t *testing.T) {

	var e *Engine

	BeforeEach(func() {
		e = Start(true)
	})

	AfterEach(func() {
		e.Stop()
	})

	Describe("NewPinNumber()", func() {
		It("should return an instance of PinIntput", func() {
			AssertEqual(reflect.TypeOf(NewPin(1)).String(), "*engine.Pin")
		})
		It("should return a rpio.Pin", func() {
			in := NewPin(1)
			AssertEqual(reflect.TypeOf(in.PinNumber()).String(), "int")
		})
	})

	Describe("Name() and String()", func() {
		It("should return a value of Pin", func() {
			in := NewPin(1)
			AssertEqual(in.String(), "Pin")
		})
		It("should return a value of FooBar", func() {
			in := NewPin(1)
			in.Name("FooBar")
			AssertEqual(in.String(), "FooBar")
		})
	})

	Report(t)
}
