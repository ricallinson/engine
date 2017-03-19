package rpio

import (
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestRpio(t *testing.T) {

	BeforeEach(func() {
		Mock = true
	})

	AfterEach(func() {
		Mock = false
	})

	Describe("rpio.Open()", func() {
		It("should return NOT an error", func() {
			err := Open()
			AssertEqual(err == nil, true)
		})
		It("should return an error", func() {
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
			Mock = false
			err := Close()
			AssertEqual(err != nil, true)
		})
	})

	Describe("rpio.PinMode() and rpio.MockGetPinMode()", func() {
		It("should set pin to Input", func() {
			PinMode(1, Input)
			AssertEqual(MockGetPinMode(1), Input)
		})
		It("should set pin to Output", func() {
			PinMode(1, Output)
			AssertEqual(MockGetPinMode(1), Output)
		})
		It("should return zero", func() {
			AssertEqual(MockGetPinMode(2), Input)
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

	Describe("rpio.PullMode() and MockGetPullMode()", func() {
		It("should set pin to PullOff", func() {
			PullMode(1, PullOff)
			AssertEqual(MockGetPullMode(1), PullOff)
		})
		It("should set pin to PullUp", func() {
			PullMode(1, PullUp)
			AssertEqual(MockGetPullMode(1), PullUp)
		})
		It("should set pin to PullDown", func() {
			PullMode(1, PullDown)
			AssertEqual(MockGetPullMode(1), PullDown)
		})
		It("should return zero", func() {
			AssertEqual(MockGetPullMode(2), PullOff)
		})
	})

	Report(t)
}
