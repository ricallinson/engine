//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package main

import (
	"github.com/ricallinson/engine"
	"time"
)

func main() {

	// Start the robots processing engine.
	var robot = engine.Start(false)

	// Create a new instance of a Motor for the left wheel.
	var motorLeft = robot.NewMotor(6, 13, 19, false)

	// Create a new instance of a Motor for the right wheel.
	var motorRight = robot.NewMotor(16, 20, 21, true)

	// Create a new instance of an IRSensor for the left of the line.
	var lineSensorLeft = robot.NewIRSensor(8)

	// Create a new instance of an IRSensor for the right of the line.
	var lineSensorRight = robot.NewIRSensor(11)

	// Start a loop that will run endlessly.
	for {
		// This will make the robot move in a straight line.
		motorLeft.Set(0.6)
		motorRight.Set(0.6)

		// Check if the left senors has a value.
		if lineSensorLeft.Get() > 0 {
			// If it does then the robot needs to move to the left.
			motorRight.Stop()
		}

		// Check if the left senors has a value.
		if lineSensorRight.Get() > 0 {
			// If it does then the robot needs to move to the right.
			motorLeft.Stop()
		}

		// All the above happen in less than 1,000s of a second.
		// To give the robot time to move we put the code to sleep for 1/4 of second.
		time.Sleep(time.Second / 200)
	}

	// At the end of the program it's good practice to stop the processing engine.
	robot.Stop()
}
