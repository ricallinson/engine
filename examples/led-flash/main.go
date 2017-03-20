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

	// Create a new instance of a LED.
	var led = robot.NewLED(2)

	// Run a loop for 20 times, each time toggling the LED from on to off, or off to on.
	for x := 0; x < 20; x++ {
		led.Toggle()
		time.Sleep(time.Second / 5)
	}

	// At the end of the program it's good practice to stop the processing engine.
	robot.Stop()
}
