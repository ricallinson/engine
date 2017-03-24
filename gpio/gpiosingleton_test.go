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

	BeforeEach(func() {

	})

	AfterEach(func() {

	})

	Describe("Gpio()", func() {
		It("should return an instance of gpioSingleton", func() {
			AssertEqual(reflect.TypeOf(Gpio()).String(), "*gpio.gpioSingleton")
		})
	})

	Report(t)
}
