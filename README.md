# Engine

[![Build Status](https://travis-ci.org/ricallinson/engine.svg?branch=master)](https://travis-ci.org/ricallinson/engine)

The purpose of this package is to programmatically control the GPIO pins on a Raspberry Pi. In doing so children and adults of all ages can make crazy robots that do all manner of crazy things.

All the parts to make a robot that runs the code in this package can be ordered from amazon for under a $100.

* [Raspberry Pi 3 Model B](https://www.amazon.com/gp/product/B01EW3QU22/ref=oh_aui_detailpage_o02_s00?ie=UTF8&psc=1)
* [Robot Car Chassis](https://www.amazon.com/gp/product/B01LXY7CM3/ref=oh_aui_detailpage_o04_s00?ie=UTF8&psc=1)
* [L293D Stepper Motor Driver](https://www.amazon.com/gp/product/B00ODQM8KC/ref=oh_aui_detailpage_o03_s00?ie=UTF8&psc=1)
* [IR Infrared Obstacle Avoidance Sensor](https://www.amazon.com/gp/product/B01I57HIJ0/ref=oh_aui_detailpage_o04_s00?ie=UTF8&psc=1)
* [Solderless BreadBoard](https://www.amazon.com/gp/product/B01258UZMC/ref=oh_aui_detailpage_o01_s00?ie=UTF8&psc=1)
* [Jumper Wires](https://www.amazon.com/gp/product/B01EV70C78/ref=oh_aui_detailpage_o01_s00?ie=UTF8&psc=1)

__Unstable__: This package is under development.

## Documentation

All the source code is commented and also available as [online documentation](https://godoc.org/github.com/ricallinson/engine).

There are working code examples of each supported controllable device.

* [LED](https://github.com/ricallinson/engine/blob/master/examples/led-flash/main.go)
* [IRSensor](https://github.com/ricallinson/engine/blob/master/examples/ir-sensor/main.go)
* [Motor](https://github.com/ricallinson/engine/blob/master/examples/motor/main.go)

A complete program for a basic line following robot can be found [here](https://github.com/ricallinson/engine/blob/master/examples/line-follower/main.go).

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

[Wiring diagram](https://github.com/ricallinson/engine/blob/master/examples/led-flash/led-flash_bb.png)

	cd ~/robot/src/github.com/ricallinson/engine/examples/led-flash
	go install
	led-flash

### IRSensor

[Wiring diagram](https://github.com/ricallinson/engine/blob/master/examples/ir-sensor/ir-sensor_bb.png)

	cd ~/robot/src/github.com/ricallinson/engine/examples/ir-sensor
	go install
	ir-sensor

### Motor

![Wiring diagram](https://raw.githubusercontent.com/ricallinson/engine/master/examples/motor/motor_bb.png)

	cd ~/robot/src/github.com/ricallinson/engine/examples/motor
	go install
	motor

### Line Follower

	cd ~/robot/src/github.com/ricallinson/engine/examples/line-follower
	go install
	line-follower

## Generate Code Coverage from Engine Tests

Run a packages tests and generate its coverage report with a HTML viewer;

	cd ~/robot/src/github.com/ricallinson/engine
	go test -coverprofile=./coverage.out; go tool cover -html=./coverage.out -o=./coverage.html
	open ./coverage.html
