package rpio

import (
	. "github.com/ricallinson/simplebdd"
	"runtime"
	"testing"
	"time"
)

func TestRpio(t *testing.T) {

	BeforeEach(func() {
		Mock = runtime.GOARCH != "arm"
		Open()
	})

	AfterEach(func() {
		Close()
		Mock = false
	})

	Describe("rpio.Open()", func() {
		It("should return NOT an error", func() {
			err := Open()
			AssertEqual(err == nil, true)
		})
		It("should return an error", func() {
			if !Mock {
				return
			}
			Mock = false
			err := Open()
			AssertEqual(err != nil, true)
		})
	})

	Describe("rpio.Close()", func() {
		It("should return NOT an error", func() {
			err := Close()
			AssertEqual(err == nil, true)
		})
		It("should return an error", func() {
			if !Mock {
				return
			}
			Mock = false
			err := Close()
			AssertEqual(err != nil, true)
		})
	})

	Describe("rpio.PinMode() and rpio.StoredPinMode()", func() {
		It("should set pin to Input", func() {
			PinMode(1, Input)
			AssertEqual(StoredPinMode(1), Input)
		})
		It("should set pin to Output", func() {
			PinMode(1, Output)
			AssertEqual(StoredPinMode(1), Output)
		})
		It("should return zero", func() {
			AssertEqual(StoredPinMode(2), Input)
		})
	})

	Describe("rpio.WritePin() and rpio.ReadPin()", func() {
		It("should set pin to High", func() {
			WritePin(1, High)
			AssertEqual(ReadPin(1), High)
		})
		It("should set pin to High", func() {
			WritePin(1, Low)
			AssertEqual(ReadPin(1), Low)
		})
	})

	Describe("rpio.WritePinPWM() and rpio.StoredPinPWM()", func() {
		It("should set PWM to 0 and call pin.Low()", func() {
			WritePinPWM(1, 0)
			AssertEqual(ReadPin(1), Low)
		})
		It("should set PWM to 0 and call pin.High()", func() {
			WritePinPWM(1, 100)
			AssertEqual(ReadPin(1), High)
		})
		It("should set pin PWM percentage to 50%", func() {
			WritePinPWM(1, 50)
			AssertEqual(StoredPinPWM(1), 50)
		})
		It("should set PWM to 0 and call pin.Low() from -0.1", func() {
			WritePinPWM(1, int(float32(-0.1)*100))
			AssertEqual(ReadPin(1), Low)
		})
	})

	Describe("rpio.PullMode() and StoredPullMode()", func() {
		It("should set pin to PullOff", func() {
			PullMode(1, PullOff)
			AssertEqual(StoredPullMode(1), PullOff)
		})
		It("should set pin to PullUp", func() {
			PullMode(1, PullUp)
			AssertEqual(StoredPullMode(1), PullUp)
		})
		It("should set pin to PullDown", func() {
			PullMode(1, PullDown)
			AssertEqual(StoredPullMode(1), PullDown)
		})
		It("should return zero", func() {
			AssertEqual(StoredPullMode(2), PullOff)
		})
	})

	Describe("rpio.Pin", func() {
		It("should set the pin to Input", func() {
			p := Pin(1)
			p.Input()
			AssertEqual(StoredPinMode(p), Input)
		})
		It("should set the pin to Output", func() {
			p := Pin(1)
			p.Output()
			AssertEqual(StoredPinMode(p), Output)
		})
		It("should set the pin to High", func() {
			p := Pin(1)
			p.High()
			AssertEqual(ReadPin(p), High)
		})
		It("should set the pin to Low", func() {
			p := Pin(1)
			p.Low()
			AssertEqual(ReadPin(p), Low)
		})
		It("should Toggle the pin state", func() {
			p := Pin(1)
			p.Toggle()
			AssertEqual(ReadPin(p), High)
			p.Toggle()
			AssertEqual(ReadPin(p), Low)
		})
		It("should set the pin mode to Input and then Output", func() {
			p := Pin(1)
			p.Mode(Input)
			AssertEqual(StoredPinMode(p), Input)
			p.Mode(Output)
			AssertEqual(StoredPinMode(p), Output)
		})
		It("should set the pin state to High and then Low", func() {
			p := Pin(1)
			p.Write(High)
			AssertEqual(ReadPin(p), High)
			p.Write(Low)
			AssertEqual(ReadPin(p), Low)
		})
		It("should read the pin state as High and then Low", func() {
			p := Pin(1)
			p.Write(High)
			AssertEqual(p.Read(), High)
			p.Write(Low)
			AssertEqual(p.Read(), Low)
		})
		It("should read the PWM pin state as High and then Low", func() {
			p := Pin(1)
			p.WritePWM(100)
			AssertEqual(p.Read(), High)
			p.WritePWM(0)
			AssertEqual(p.Read(), Low)
		})
		It("should read the PWM pin percentage as 50%", func() {
			p := Pin(1)
			p.WritePWM(50)
			AssertEqual(StoredPinPWM(1), 50)
		})
		It("should read the PWM pin percentage as increasing from 0% to 100%", func() {
			p := Pin(1)
			brightness := 0
			for x := 0; x < 10; x++ {
				p.WritePWM(brightness)
				AssertEqual(StoredPinPWM(1), brightness)
				brightness = brightness + 10
				time.Sleep(10 * time.Millisecond)
			}
		})
		It("should set and read the pin state to High, 50% then Low each time with the same pin", func() {
			p := Pin(1)
			for x := 0; x < 10; x++ {
				p.High()
				AssertEqual(ReadPin(1), High)
				p.WritePWM(50)
				p.Low()
				AssertEqual(ReadPin(1), Low)
			}
		})
		It("should set and read the pin state to High, 50% then Low each time with a new pin", func() {
			for x := 0; x < 10; x++ {
				p := Pin(1)
				p.High()
				AssertEqual(ReadPin(1), High)
				p.WritePWM(50)
				p.Low()
				AssertEqual(ReadPin(1), Low)
			}
		})
		It("should set and read the pin state to High, 50% then Low from a new instance of rpio each time", func() {
			for x := 0; x < 10; x++ {
				p := Pin(1)
				p.High()
				AssertEqual(ReadPin(1), High)
				p.WritePWM(50)
				p.Low()
				AssertEqual(ReadPin(1), Low)
				Close()
				Open()
			}
		})
		It("should set the pin pull to PullUp, PullDown, PullOff", func() {
			p := Pin(1)
			p.Pull(PullUp)
			AssertEqual(StoredPullMode(p), PullUp)
			p.Pull(PullDown)
			AssertEqual(StoredPullMode(p), PullDown)
			p.Pull(PullOff)
			AssertEqual(StoredPullMode(p), PullOff)
		})
		It("should set the pin pull to PullUp, PullDown, PullOff using functions", func() {
			p := Pin(1)
			p.PullUp()
			AssertEqual(StoredPullMode(p), PullUp)
			p.PullDown()
			AssertEqual(StoredPullMode(p), PullDown)
			p.PullOff()
			AssertEqual(StoredPullMode(p), PullOff)
		})
	})

	Report(t)
}
