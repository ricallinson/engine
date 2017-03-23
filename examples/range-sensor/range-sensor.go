//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package main

import (
	"github.com/ricallinson/engine"
	"log"
	"time"
)

func main() {

	// Start the robots processing engine.
	var robot = engine.Start(false)

	// Create a new instance of a rangeSensor.
	var rangeSensor = robot.NewRangeSensor(23, 24)

	// Run a loop for 20 times, each time getting the current measured range.
	for x := 0; x < 20; x++ {
		log.Println("Measuring distance in cm...")
		cm := rangeSensor.Get()
		if cm < 0 {
			log.Println("Range measurement failed.")
		}
		time.Sleep(time.Second)
	}

	// At the end of the program it's good practice to stop the processing engine.
	robot.Stop()
}
