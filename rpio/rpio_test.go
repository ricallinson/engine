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

	Report(t)
}
