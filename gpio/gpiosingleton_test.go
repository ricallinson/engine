//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package gpio

import (
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func TestGpioSingleton(t *testing.T) {

	var gpio *GpioSingleton

	BeforeEach(func() {
		gpio = Gpio()
	})

	AfterEach(func() {
		gpio.Close()
	})

	Describe("Gpio()", func() {
		It("should return an instance of GpioSingleton", func() {
			AssertEqual(reflect.TypeOf(gpio).String(), "*gpio.GpioSingleton")
		})
	})

	Describe("gpio.open()", func() {
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

	Describe("gpio.IsMock()", func() {
		It("should return NOT an error", func() {
			AssertEqual(gpio.IsMock(), runtime.GOARCH != "arm")
		})
	})

	Describe("gpio.PinOut()", func() {
		It("should return an instance of GpioPin", func() {
			AssertEqual(reflect.TypeOf(gpio.Pin(1)).String(), "*gpio.GpioPin")
		})
		It("should get the same pin twice", func() {
			pA := gpio.Pin(1)
			pB := gpio.Pin(1)
			pC := gpio.Pin(2)
			AssertEqual(pA, pB)
			AssertNotEqual(pA, pC)
		})
		It("should set the pin to Input", func() {
			p := gpio.Pin(1)
			p.Input()
			AssertEqual(p.GetMode(), Input)
		})
		It("should set the pin to Output", func() {
			p := gpio.Pin(1)
			p.Output()
			AssertEqual(p.GetMode(), Output)
		})
		It("should set the pin to High", func() {
			p := gpio.Pin(1)
			p.High()
			AssertEqual(p.Read(), High)
		})
		It("should set the pin to Low", func() {
			p := gpio.Pin(1)
			p.Low()
			AssertEqual(p.Read(), Low)
		})
		It("should Toggle the pin state", func() {
			p := gpio.Pin(1)
			p.Toggle()
			AssertEqual(p.Read(), High)
			p.Toggle()
			AssertEqual(p.Read(), Low)
		})
		It("should set the pin mode to Input and then Output", func() {
			p := gpio.Pin(1)
			p.Mode(Input)
			AssertEqual(p.GetMode(), Input)
			p.Mode(Output)
			AssertEqual(p.GetMode(), Output)
		})
		It("should set the pin state to High and then Low", func() {
			p := gpio.Pin(1)
			p.Write(High)
			AssertEqual(p.Read(), High)
			p.Write(Low)
			AssertEqual(p.Read(), Low)
		})
		It("should read the pin state as High and then Low", func() {
			p := gpio.Pin(1)
			p.Write(High)
			AssertEqual(p.Read(), High)
			p.Write(Low)
			AssertEqual(p.Read(), Low)
		})
		It("should read the PWM pin state as High and then Low", func() {
			p := gpio.Pin(1)
			p.Modulate(100, 100)
			AssertEqual(p.Read(), High)
			p.Modulate(0, 100)
			AssertEqual(p.Read(), Low)
		})
		It("should read the PWM pin percentage as 50%", func() {
			p := gpio.Pin(1)
			p.Modulate(50, 100)
			AssertEqual(p.GetModulation(), 50)
		})
		It("should read the PWM pin percentage as increasing from 0% to 100%", func() {
			p := gpio.Pin(1)
			brightness := 0
			for x := 0; x < 10; x++ {
				p.Modulate(brightness, 100)
				AssertEqual(p.GetModulation(), brightness)
				brightness = brightness + 10
				time.Sleep(10 * time.Millisecond)
			}
		})
		It("should set and read the pin state to High, 50% then Low each time with the same pin", func() {
			p := gpio.Pin(1)
			for x := 0; x < 10; x++ {
				p.High()
				AssertEqual(p.Read(), High)
				p.Modulate(50, 100)
				p.Low()
				AssertEqual(p.Read(), Low)
			}
		})
		It("should set and read the pin state to High, 50% then Low from a new instance of rpio each time", func() {
			for x := 0; x < 10; x++ {
				p := gpio.Pin(1)
				p.High()
				AssertEqual(p.Read(), High)
				p.Modulate(50, 100)
				p.Low()
				AssertEqual(p.Read(), Low)
				gpio.Close()
				gpio = Gpio()
			}
		})
		It("should set the pin pull to PullUp, PullDown, PullOff", func() {
			p := gpio.Pin(1)
			p.Pull(PullUp)
			AssertEqual(p.GetPull(), PullUp)
			p.Pull(PullDown)
			AssertEqual(p.GetPull(), PullDown)
			p.Pull(PullOff)
			AssertEqual(p.GetPull(), PullOff)
		})
		It("should set the pin pull to PullUp, PullDown, PullOff using functions", func() {
			p := gpio.Pin(1)
			p.PullUp()
			AssertEqual(p.GetPull(), PullUp)
			p.PullDown()
			AssertEqual(p.GetPull(), PullDown)
			p.PullOff()
			AssertEqual(p.GetPull(), PullOff)
		})
	})

	Report(t)
}
