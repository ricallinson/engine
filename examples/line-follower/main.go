package main

import (
	"github.com/ricallinson/engine"
	"time"
)

func main() {

	// Start the robots processing engine.
	var robot = engine.Start()

	// Create a new instance of a Motor for the left wheel.
	var motorLeft = robot.NewMotor(1, true)

	// Create a new instance of a Motor for the right wheel.
	var motorRight = robot.NewMotor(2, false)

	// Create a new instance of an IRSensor for the left of the line.
	var lineSensorLeft = robot.NewIRSensor(3)

	// Create a new instance of an IRSensor for the right of the line.
	var lineSensorRight = robot.NewIRSensor(4)

	// Start a loop that will run endlessly.
	for {
		// Set the power for the left and right motors.
		// This will make the robot move in a straight line.
		motorLeft.Set(1)
		motorRight.Set(1)

		// Check if the left senors has a value.
		if lineSensorLeft.Get() > 0 {
			// If it does then the robot needs to move to the left.
			// Set the power to 0 to stop the left motor.
			motorLeft.Set(0)
		}

		// Check if the left senors has a value.
		if lineSensorRight.Get() > 0 {
			// If it does then the robot needs to move to the right.
			// Set the power to 0 to stop the right motor.
			motorRight.Set(0)
		}

		// All the above happen in less than 1,000s of a second.
		// To give the robot time to move we put the code to sleep for 1/4 of second.
		time.Sleep(time.Second / 4)
	}

	// At the end of the program it's good practice to stop the processing engine.
	robot.Stop()
}
