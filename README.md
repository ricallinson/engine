# Engine

[![Build Status](https://travis-ci.org/ricallinson/engine.svg?branch=master)](https://travis-ci.org/ricallinson/engine)

The purpose of this package is to programmatically control the GPIO pins on a Raspberry Pi. In doing so children and adults of all ages can make crazy robots that do all manner of crazy things.

All the parts to make a robot that runs the code in this package can be ordered from amazon for under a $100;

* [Raspberry Pi 3 Model B](https://www.amazon.com/gp/product/B01EW3QU22/ref=oh_aui_detailpage_o02_s00?ie=UTF8&psc=1)
* [Robot Car Chassis](https://www.amazon.com/gp/product/B01LXY7CM3/ref=oh_aui_detailpage_o04_s00?ie=UTF8&psc=1)
* [L293D Stepper Motor Driver](https://www.amazon.com/gp/product/B00ODQM8KC/ref=oh_aui_detailpage_o03_s00?ie=UTF8&psc=1)
* [IR Infrared Obstacle Avoidance Sensor](https://www.amazon.com/gp/product/B01I57HIJ0/ref=oh_aui_detailpage_o04_s00?ie=UTF8&psc=1)
* [Solderless BreadBoard](https://www.amazon.com/gp/product/B01258UZMC/ref=oh_aui_detailpage_o01_s00?ie=UTF8&psc=1)
* [Jumper Wires](https://www.amazon.com/gp/product/B01EV70C78/ref=oh_aui_detailpage_o01_s00?ie=UTF8&psc=1)

Optional parts for more fun;

* [HC-SR04 Ultrasonic Distance Sensor](https://www.amazon.com/Elegoo-HC-SR04-Ultrasonic-Distance-MEGA2560/dp/B01COSN7O6/ref=sr_1_5?ie=UTF8&qid=1490247271&sr=8-5&keywords=hc-sr04)
* [10 Ohm - 1M Ohm Resistor Pack](https://www.amazon.com/E-Projects-EPC-103-Value-Resistor-Kit/dp/B00E9YQQSS/ref=sr_1_3?ie=UTF8&qid=1490247363&sr=8-3&keywords=resistor)

__Unstable__: This package is under development.

## Documentation

All the source code is commented and also available as [online documentation](https://godoc.org/github.com/ricallinson/engine).

There are working code examples of each supported controllable device.

* [LED](#led)
* [IRSensor](#irsensor)
* [Motor](#motor)

A complete program for a basic line following robot can be found [here](#line-follower).

## Working Notes

### Setup the Raspberry Pi

You will need to create a SSD with [Raspbian Jesse Lite](https://www.raspberrypi.org/downloads/raspbian/). Follow the Raspberry Pi instructions on how to [install an operating system](https://www.raspberrypi.org/documentation/installation/installing-images/README.md).

It's a good idea to use a [Secure Shell](https://www.raspberrypi.org/documentation/remote-access/ssh/) when working with your robot but not required.

### Setup the Environment

After logging on to the Raspberry Pi execute the following commands;

	sudo apt-get install git
	sudo apt-get install golang
	mkdir ~/robot
	cd ~/robot

Add the environment variables for using Go;
	
	export PATH=$PATH:$GOROOT/bin
	export GOPATH=$HOME/robot
	export PATH=$PATH:$GOPATH/bin

Get the [Engine](https://github.com/ricallinson/engine) from Github;

	go get github.com/ricallinson/simplebdd
	go get github.com/ricallinson/engine

### Subsequent Logins or Refreshing the Environment

	export PATH=$PATH:$GOROOT/bin
	export GOPATH=$HOME/robot
	export PATH=$PATH:$GOPATH/bin
	cd ~/robot
	go get -u github.com/ricallinson/engine

## Testing

### Run Engine Tests

	cd ~/robot/src/github.com/ricallinson/engine
	go test -cover ./...

## Examples

### LED

Source code for [LED](https://github.com/ricallinson/engine/blob/master/examples/led-flash/main.go) exmaple.

![Wiring diagram](https://raw.githubusercontent.com/ricallinson/engine/master/examples/led-flash/led-flash_bb.png)

	cd ~/robot/src/github.com/ricallinson/engine/examples/led-flash
	go install
	led-flash

### IRSensor

Source code for [IRSensor](https://github.com/ricallinson/engine/blob/master/examples/ir-sensor/main.go) example.

![Wiring diagram](https://raw.githubusercontent.com/ricallinson/engine/master/examples/ir-sensor/ir-sensor_bb.png)

	cd ~/robot/src/github.com/ricallinson/engine/examples/ir-sensor
	go install
	ir-sensor

### RangeSensor

Source code for [RangeSensor](https://github.com/ricallinson/engine/blob/master/examples/range-sensor/main.go) example. Requires optional parts to complete.

![Wiring diagram](https://raw.githubusercontent.com/ricallinson/engine/master/examples/range-sensor/range-sensor_bb.png)

	cd ~/robot/src/github.com/ricallinson/engine/examples/range-sensor
	go install
	range-sensor

### Motor

Source code for [Motor](https://github.com/ricallinson/engine/blob/master/examples/motor/main.go) example.

![Wiring diagram](https://raw.githubusercontent.com/ricallinson/engine/master/examples/motor/motor_bb.png)

	cd ~/robot/src/github.com/ricallinson/engine/examples/motor
	go install
	motor

### Line Follower

Source code for [Line Follower](https://github.com/ricallinson/engine/blob/master/examples/line-follower/main.go) example.

![Wiring diagram](https://raw.githubusercontent.com/ricallinson/engine/master/examples/line-follower/line-follower_bb.png)

	cd ~/robot/src/github.com/ricallinson/engine/examples/line-follower
	go install
	line-follower

## Generate Code Coverage from Engine Tests

Run a packages tests and generate its coverage report with a HTML viewer;

	cd ~/robot/src/github.com/ricallinson/engine
	go test -coverprofile=./coverage.out; go tool cover -html=./coverage.out -o=./coverage.html
	open ./coverage.html
