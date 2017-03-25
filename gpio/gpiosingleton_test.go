//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package gpio

import (
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
)

func TestGpioSingleton(t *testing.T) {

	var gpio *gpioSingleton

	BeforeEach(func() {
		gpio = Gpio()
	})

	AfterEach(func() {
		gpio.Close()
	})

	Describe("Gpio()", func() {
		It("should return an instance of gpioSingleton", func() {
			AssertEqual(reflect.TypeOf(gpio).String(), "*gpio.gpioSingleton")
		})
	})

	Describe("gpio.Open()", func() {
		It("should return NOT an error", func() {
			err := Gpio().open()
			AssertEqual(err == nil, true)
		})
		It("should return an error", func() {
			gpio.mock = false
			err := Gpio().open()
			AssertEqual(err != nil, true)
		})
	})

	Describe("gpio.Close()", func() {
		It("should return NOT an error", func() {
			err := gpio.Close()
			AssertEqual(err == nil, true)
		})
		It("should return an error", func() {
			gpio.mock = false
			err := gpio.Close()
			AssertEqual(err != nil, true)
		})
	})

	Describe("gpio.Pin()", func() {
		It("should return an instance of Pin", func() {
			AssertEqual(reflect.TypeOf(gpio.Pin(1)).String(), "*gpio.Pin")
		})
	})

	Report(t)
}
