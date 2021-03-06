//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package engine

import (
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"runtime"
	"testing"
)

func TestEngine(t *testing.T) {

	var e *Engine

	BeforeEach(func() {
		e = Start(runtime.GOARCH != "arm")
	})

	AfterEach(func() {
		e.Stop()
	})

	Describe("NewEngine()", func() {
		It("should return an instance of Engine", func() {
			AssertEqual(reflect.TypeOf(e).String(), "*engine.Engine")
		})
	})

	Describe("NewToggleSwitch()", func() {
		It("should return an instance of ToggleSwitch", func() {
			AssertEqual(reflect.TypeOf(e.NewToggleSwitch(1)).String(), "*engine.ToggleSwitch")
		})
		It("should fail as pin has alreay been used", func() {
			defer func() {
				AssertEqual(recover() != nil, true)
			}()
			e.NewToggleSwitch(1)
			e.NewToggleSwitch(1)
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
			AssertEqual(reflect.TypeOf(e.NewMotor(1, 2, 3, false)).String(), "*engine.Motor")
		})
		It("should return an instance of Motor with direction reversed", func() {
			AssertEqual(reflect.TypeOf(e.NewMotor(1, 2, 3, true)).String(), "*engine.Motor")
		})
		It("should fail as pin has alreay been used", func() {
			defer func() {
				AssertEqual(recover() != nil, true)
			}()
			e.NewMotor(1, 2, 3, true)
			e.NewMotor(3, 2, 1, true)
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

	Describe("NewRangeSensor()", func() {
		It("should return an instance of RangeSensor", func() {
			AssertEqual(reflect.TypeOf(e.NewRangeSensor(1, 2)).String(), "*engine.RangeSensor")
		})
		It("should fail as pin has alreay been used", func() {
			defer func() {
				AssertEqual(recover() != nil, true)
			}()
			e.NewRangeSensor(1, 2)
			e.NewRangeSensor(2, 1)
		})
	})

	Describe("registerPinNumber()", func() {
		It("should NOT panic as pin 1 is in range", func() {
			defer func() {
				AssertEqual(recover() == nil, true)
			}()
			e.registerPin(1)
		})
		It("should NOT panic as pin 25 is in range", func() {
			defer func() {
				AssertEqual(recover() == nil, true)
			}()
			e.registerPin(25)
		})
		It("should panic as pin 26 is upward of range", func() {
			defer func() {
				AssertEqual(recover() != nil, true)
			}()
			e.registerPin(26)
		})
		It("should panic as pin 0 is lower than range", func() {
			defer func() {
				AssertEqual(recover() != nil, true)
			}()
			e.registerPin(0)
		})
	})

	Report(t)
}
