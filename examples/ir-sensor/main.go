package main

import (
	"github.com/ricallinson/engine"
	"time"
)

func main() {

	// Start the robots processing engine.
	var robot = engine.Start(false)

	// Create a new instance of a IRSensor.
	var ir = robot.NewIRSensor(8)

	// Run a loop for 20 times, each time getting the current value from the IRSensor.
	for x := 0; x < 20; x++ {
		ir.Get()
		// Sleep for 1 second in between getting a value.
		time.Sleep(time.Second)
	}

	// At the end of the program it's good practice to stop the processing engine.
	robot.Stop()
}
