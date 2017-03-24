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

	// Create a variable to hold the brightness percentage.
	var brightness float32 = 0.10

	// Run a loop for 10 times, each time increasing the brightness of the LED by 10%.
	for x := 0; x < 10; x++ {
		led.Set(brightness)
		brightness = brightness + 0.10
		time.Sleep(time.Second * 2)
	}

	// Turn the LED off before exiting the program.
	led.Off()

	// At the end of the program it's good practice to stop the processing engine.
	robot.Stop()
}
