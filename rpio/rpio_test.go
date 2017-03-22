package rpio

import (
	. "github.com/ricallinson/simplebdd"
	"runtime"
	"testing"
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

	Describe("rpio.WritePin() rpio.ReadPin()", func() {
		It("should set pin to High", func() {
			WritePin(1, High)
			AssertEqual(ReadPin(1), High)
		})
		It("should set pin to High", func() {
			WritePin(1, Low)
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
		It("should read the pin state to High and then Low", func() {
			p := Pin(1)
			p.Write(High)
			AssertEqual(p.Read(), High)
			p.Write(Low)
			AssertEqual(p.Read(), Low)
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
