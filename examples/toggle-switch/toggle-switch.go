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

	// Create a new instance of a ToggleSwitch.
	var toggleSwitch = robot.NewToggleSwitch(2)

	// Run in a loop until the toggle switch is turn on.
	for toggleSwitch.Get() < 1 {
		time.Sleep(time.Second / 5)
	}

	// At the end of the program it's good practice to stop the processing engine.
	robot.Stop()
}
