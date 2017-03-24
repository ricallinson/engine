/*
Extended from https://github.com/stianeikeland/go-rpio

Package rpio provides GPIO access on the Raspberry PI without any need
for external c libraries (ex: WiringPI or BCM2835).

Supports simple operations such as:
- Pin mode/direction (input/output)
- Pin write (high/low)
- Pin read (high/low)
- Pull up/down/off

Example of use:

	rpio.Open()
	defer rpio.Close()

	pin := rpio.Pin(4)
	pin.Output()

	for {
		pin.Toggle()
		time.Sleep(time.Second)
	}

The library use the raw BCM2835 pinouts, not the ports as they are mapped
on the output pins for the raspberry pi

   Rev 1 Raspberry Pi
+------+------+--------+
| GPIO | Phys | Name   |
+------+------+--------+
|   0  |   3  | SDA    |
|   1  |   5  | SCL    |
|   4  |   7  | GPIO 7 |
|   7  |  26  | CE1    |
|   8  |  24  | CE0    |
|   9  |  21  | MISO   |
|  10  |  19  | MOSI   |
|  11  |  23  | SCLK   |
|  14  |   8  | TxD    |
|  15  |  10  | RxD    |
|  17  |  11  | GPIO 0 |
|  18  |  12  | GPIO 1 |
|  21  |  13  | GPIO 2 |
|  22  |  15  | GPIO 3 |
|  23  |  16  | GPIO 4 |
|  24  |  18  | GPIO 5 |
|  25  |  22  | GPIO 6 |
+------+------+--------+

See the spec for full details of the BCM2835 controller:
http://www.raspberrypi.org/wp-content/uploads/2012/02/BCM2835-ARM-Peripherals.pdf

*/

package rpio

import (
	"bytes"
	"encoding/binary"
	"os"
	"reflect"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

type Direction uint8
type Pin uint8
type State uint8
type Pull uint8

// Memory offsets for gpio, see the spec for more details
const (
	bcm2835Base = 0x20000000
	pi1GPIOBase = bcm2835Base + 0x200000
	memLength   = 4096

	pinMask uint32 = 7 // 0b111 - pinmode is 3 bits
)

// Pin direction, a pin can be set in Input or Output mode
const (
	Input Direction = iota
	Output
)

// State of pin, High / Low
const (
	Low State = iota
	High
)

// Pull Up / Down / Off
const (
	PullOff Pull = iota
	PullDown
	PullUp
)

// General constants
const (
	PinCount = 26
)

// Arrays for 8 / 32 bit access to memory and a semaphore for write locking
var (
	memlock sync.Mutex
	mem     []uint32
	mem8    []uint8
)

// Arrays for testing and a flag for using a rpio as a Mock.
var (
	storePinPWM       []int
	storePinPull      []Pull
	storePinState     []State
	storePinDiredtion []Direction
	Mock              bool
)

// Set pin as Input
func (pin Pin) Input() {
	PinMode(pin, Input)
}

// Set pin as Output
func (pin Pin) Output() {
	PinMode(pin, Output)
}

// Set pin High
func (pin Pin) High() {
	WritePin(pin, High)
}

// Set pin Low
func (pin Pin) Low() {
	WritePin(pin, Low)
}

// Toggle pin state
func (pin Pin) Toggle() {
	TogglePin(pin)
}

// Set pin Direction
func (pin Pin) Mode(dir Direction) {
	PinMode(pin, dir)
}

// Set pin state (high/low)
func (pin Pin) Write(state State) {
	WritePin(pin, state)
}

// Takes a range from 0% to 100% as an integer and sets the pin.High() to pulse at that percentage of 1/500 of a second.
func (pin Pin) WritePWM(pwm int) {
	WritePinPWM(pin, pwm)
}

// Read pin state (high/low)
func (pin Pin) Read() State {
	return ReadPin(pin)
}

// Set a given pull up/down mode
func (pin Pin) Pull(pull Pull) {
	PullMode(pin, pull)
}

// Pull up pin
func (pin Pin) PullUp() {
	PullMode(pin, PullUp)
}

// Pull down pin
func (pin Pin) PullDown() {
	PullMode(pin, PullDown)
}

// Disable pullup/down on pin
func (pin Pin) PullOff() {
	PullMode(pin, PullOff)
}

// PinMode sets the direction of a given pin (Input or Output)
func PinMode(pin Pin, direction Direction) {
	// Store the pin direction for testing.
	storePinDiredtion[pin] = direction

	// If the Mock flag is set do nothing.
	if Mock {
		return
	}

	// Pin fsel register, 0 or 1 depending on bank
	fsel := uint8(pin) / 10
	shift := (uint8(pin) % 10) * 3

	memlock.Lock()
	defer memlock.Unlock()

	if direction == Input {
		mem[fsel] = mem[fsel] &^ (pinMask << shift)
	} else {
		mem[fsel] = (mem[fsel] &^ (pinMask << shift)) | (1 << shift)
	}

}

// Return the last stored mode for the give pin.
func StoredPinMode(pin Pin) Direction {
	return storePinDiredtion[pin]
}

// WritePin sets a given pin High or Low
// by setting the clear or set registers respectively
func WritePin(pin Pin, state State) {
	// If the Mock flag is return the stored value.
	if Mock {
		storePinState[pin] = state
		return
	}

	p := uint8(pin)

	// Clear register, 10 / 11 depending on bank
	// Set register, 7 / 8 depending on bank
	clearReg := p/32 + 10
	setReg := p/32 + 7

	memlock.Lock()
	defer memlock.Unlock()

	if state == Low {
		mem[clearReg] = 1 << (p & 31)
	} else {
		mem[setReg] = 1 << (p & 31)
	}
}

// Takes a range from 0% to 100% as an integer and sets the pin.High() to pulse at that percentage of 1/500 of a second.
// Loops every 2ms setting pin.High and then pin.Low for the percentage of time stored in storePinPWM[pin].
func WritePinPWM(pin Pin, modulation int) {
	// If modulation is 0 or less then reset stored modulation and call pin.Low().
	if modulation < 1 {
		storePinPWM[pin] = 0
		pin.Low()
		return
	}
	// If modulation is 100 or greater then reset stored modulation and call pin.High().
	if modulation > 99 {
		storePinPWM[pin] = 0
		pin.High()
		return
	}
	// If there is already a modulation value then update it and return.
	if storePinPWM[pin] > 0 {
		storePinPWM[pin] = modulation
		return
	}
	// If none of the above are true then store the modulation percentage for the pin.
	storePinPWM[pin] = modulation
	// Start the modulation routine that will run until the modulation value is out of range.
	go modulatePin(pin, modulation, modulations[pin])
}

// Software implemented Pulse Width Modulation (PWM) at every 2ms.
// A channel is used to stop the routine when Close() is called.
func modulatePin(pin Pin, modulation int, stage chan int) {
	// Create the int to store the High microsecond time.
	var high int
	// Started modulating on pin.
	// Check that modulation value is in range.
	for modulation > 0 && modulation < 100 {
		switch <-stage {
		case 0:
			return
		case 1:
			// The modulation is 200uS so multiply the percentage by 2.
			high = modulation * 2
			pin.High()
			// Sleep for pulse high duration.
			time.Sleep(time.Duration(high) * time.Microsecond)
			stage <- 2
		case 2:
			pin.Low()
			// Sleep for pulse low duration. The remainder of the 200uS used by pulse High.
			time.Sleep(time.Duration(200-high) * time.Microsecond)
			// Get the current PWM value for this pin.
			modulation = StoredPinPWM(pin)
			stage <- 1
		}
	}
}

// Return the last stored PWM percentage for the give pin.
func StoredPinPWM(pin Pin) int {
	return storePinPWM[pin]
}

// Read the state of a pin
func ReadPin(pin Pin) State {
	// If the Mock flag is set do nothing.
	if Mock {
		return storePinState[pin]
	}

	// Input level register offset (13 / 14 depending on bank)
	levelReg := uint8(pin)/32 + 13

	if (mem[levelReg] & (1 << uint8(pin))) != 0 {
		return High
	}

	return Low
}

// Toggle a pin state (high -> low -> high)
// TODO: probably possible to do this much faster without read
func TogglePin(pin Pin) {
	switch ReadPin(pin) {
	case Low:
		pin.High()
	case High:
		pin.Low()
	}
}

func PullMode(pin Pin, pull Pull) {
	// Store the pull mode for testing.
	storePinPull[pin] = pull

	// If the Mock flag is set do nothing.
	if Mock {
		return
	}

	// Pull up/down/off register has offset 38 / 39, pull is 37
	pullClkReg := uint8(pin)/32 + 38
	pullReg := 37
	shift := (uint8(pin) % 32)

	memlock.Lock()
	defer memlock.Unlock()

	switch pull {
	case PullDown, PullUp:
		mem[pullReg] = mem[pullReg]&^3 | uint32(pull)
	case PullOff:
		mem[pullReg] = mem[pullReg] &^ 3
	}

	// Wait for value to clock in, this is ugly, sorry :(
	time.Sleep(time.Microsecond)

	mem[pullClkReg] = 1 << shift

	// Wait for value to clock in
	time.Sleep(time.Microsecond)

	mem[pullReg] = mem[pullReg] &^ 3
	mem[pullClkReg] = 0
}

// Return the last stored pull mode for the give pin.
func StoredPullMode(pin Pin) Pull {
	return storePinPull[pin]
}

var modulations = make([]chan int, PinCount, PinCount)

// Open and memory map GPIO memory range from /dev/mem .
// Some reflection magic is used to convert it to a unsafe []uint32 pointer
func Open() (err error) {
	// Create stores.
	modulations = make([]chan int, PinCount, PinCount)
	storePinPWM = make([]int, PinCount, PinCount)
	storePinPull = make([]Pull, PinCount, PinCount)
	storePinState = make([]State, PinCount, PinCount)
	storePinDiredtion = make([]Direction, PinCount, PinCount)

	// If the Mock flag is set do nothing.
	if Mock {
		return
	}

	var file *os.File
	var base int64

	// Open fd for rw mem access; try gpiomem first
	if file, err = os.OpenFile("/dev/gpiomem", os.O_RDWR|os.O_SYNC, 0); os.IsNotExist(err) {
		file, err = os.OpenFile("/dev/mem", os.O_RDWR|os.O_SYNC, 0)
		base = getGPIOBase()
	}

	if err != nil {
		return
	}

	// FD can be closed after memory mapping
	defer file.Close()

	memlock.Lock()
	defer memlock.Unlock()

	// Memory map GPIO registers to byte array
	mem8, err = syscall.Mmap(
		int(file.Fd()),
		base,
		memLength,
		syscall.PROT_READ|syscall.PROT_WRITE,
		syscall.MAP_SHARED)

	if err != nil {
		return
	}

	// Convert mapped byte memory to unsafe []uint32 pointer, adjust length as needed
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&mem8))
	header.Len /= (32 / 8) // (32 bit = 4 bytes)
	header.Cap /= (32 / 8)

	mem = *(*[]uint32)(unsafe.Pointer(&header))

	return nil
}

// Close unmaps GPIO memory and stops all modulations.
func Close() error {
	// Stop all modulation routines.
	for _, modulation := range modulations {
		if modulation != nil {
			close(modulation)
		}
	}
	// Reset all stores.
	modulations = nil
	storePinPWM = nil
	storePinPull = nil
	storePinState = nil
	storePinDiredtion = nil

	// If the Mock flag is set do nothing.
	if Mock {
		return nil
	}

	memlock.Lock()
	defer memlock.Unlock()
	return syscall.Munmap(mem8)
}

// Read /proc/device-tree/soc/ranges and determine the base address.
// Use the default Raspberry Pi 1 base address if this fails.
func getGPIOBase() (base int64) {
	base = pi1GPIOBase
	ranges, err := os.Open("/proc/device-tree/soc/ranges")
	defer ranges.Close()
	if err != nil {
		return
	}
	b := make([]byte, 4)
	n, err := ranges.ReadAt(b, 4)
	if n != 4 || err != nil {
		return
	}
	buf := bytes.NewReader(b)
	var out uint32
	err = binary.Read(buf, binary.BigEndian, &out)
	if err != nil {
		return
	}
	return int64(out + 0x200000)
}
