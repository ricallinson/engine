package main

import (
	"github.com/ricallinson/engine"
	"time"
)

func main() {

	// Start the robots processing engine.
	var robot = engine.Start(false)

	// Create a new instance of a Motor for the left wheel.
	var motorLeft = robot.NewMotor(23, 24, 25, true)

	// Create a new instance of a Motor for the right wheel.
	var motorRight = motorLeft//robot.NewMotor(16, 20, 21, false)

	// Create a new instance of an IRSensor for the left of the line.
	var lineSensorLeft = robot.NewIRSensor(2)

	// Create a new instance of an IRSensor for the right of the line.
	var lineSensorRight = robot.NewIRSensor(3)

	// Start a loop that will run endlessly.
	for {
		// This will make the robot move in a straight line.
		motorLeft.Forwards()
		motorRight.Forwards()

		// Check if the left senors has a value.
		if lineSensorLeft.Get() > 0 {
			// If it does then the robot needs to move to the left.
			motorLeft.Stop()
		}

		// Check if the left senors has a value.
		if lineSensorRight.Get() > 0 {
			// If it does then the robot needs to move to the right.
			motorRight.Stop()
		}

		// All the above happen in less than 1,000s of a second.
		// To give the robot time to move we put the code to sleep for 1/4 of second.
		time.Sleep(time.Second / 4)
	}

	// At the end of the program it's good practice to stop the processing engine.
	robot.Stop()
}
