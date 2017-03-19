package engine

import (
	"github.com/ricallinson/engine/rpio"
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
)

func TestEngine(t *testing.T) {

	var e *Engine

	rpio.Mock = true

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
		It("should fail as Engine has alreay been started", func() {
			defer func() {
				AssertEqual(recover() != nil, true)
			}()
			Start()
		})
	})

	Describe("NewLED()", func() {
		It("should return an instance of LED", func() {
			AssertEqual(reflect.TypeOf(e.NewLED(1)).String(), "*engine.LED")
		})
		It("should fail as pin has alreay been used", func() {
			defer func() {
				AssertEqual(recover() != nil, true)
			}()
			e.NewLED(1)
			e.NewLED(1)
		})
	})

	Describe("NewMotor()", func() {
		It("should return an instance of Motor with direction forward", func() {
			AssertEqual(reflect.TypeOf(e.NewMotor(1, false)).String(), "*engine.Motor")
		})
		It("should return an instance of Motor with direction reversed", func() {
			AssertEqual(reflect.TypeOf(e.NewMotor(1, true)).String(), "*engine.Motor")
		})
		It("should fail as pin has alreay been used", func() {
			defer func() {
				AssertEqual(recover() != nil, true)
			}()
			e.NewMotor(1, true)
			e.NewMotor(1, true)
		})
	})

	Describe("NewIRSensor()", func() {
		It("should return an instance of IRSensor", func() {
			AssertEqual(reflect.TypeOf(e.NewIRSensor(1)).String(), "*engine.IRSensor")
		})
		It("should fail as pin has alreay been used", func() {
			defer func() {
				AssertEqual(recover() != nil, true)
			}()
			e.NewIRSensor(1)
			e.NewIRSensor(1)
		})
	})

	Report(t)
}
