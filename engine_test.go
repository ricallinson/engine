package engine

import (
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
)

func TestEngine(t *testing.T) {

	var e *Engine

	BeforeEach(func() {
		e = Start()
	})

	AfterEach(func() {
		e.Stop()
	})

	Describe("NewEngine()", func() {
		It("should return an instance of Engine", func() {
			AssertEqual(reflect.TypeOf(e).String(), "*engine.Engine")
		})
	})

	Report(t)
}
