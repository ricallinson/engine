//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

// Based on https://github.com/stianeikeland/go-rpio
package gpio

import (
	"bytes"
	"encoding/binary"
	"os"
	"reflect"
	"runtime"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

var GpioSingletonInstance *GpioSingleton

type GpioSingleton struct {
	pins    []*GpioPin
	mock    bool
	memlock sync.Mutex
	mem     []uint32
	mem8    []uint8
}

func Gpio() *GpioSingleton {
	if GpioSingletonInstance != nil {
		return GpioSingletonInstance
	}
	this := &GpioSingleton{
		pins: make([]*GpioPin, 26, 26),
		mock: runtime.GOARCH != "arm",
	}
	this.open()
	GpioSingletonInstance = this
	return this
}

func (this *GpioSingleton) open() error {
	// If the Mock flag is set do nothing.
	if this.mock {
		return nil
	}

	var err error
	var file *os.File
	var base int64

	// Open fd for rw mem access; try gpiomem first
	if file, err = os.OpenFile("/dev/gpiomem", os.O_RDWR|os.O_SYNC, 0); os.IsNotExist(err) {
		file, err = os.OpenFile("/dev/mem", os.O_RDWR|os.O_SYNC, 0)
		base = this.getGPIOBase()
	}

	if err != nil {
		return err
	}

	// FD can be closed after memory mapping
	defer file.Close()

	this.memlock.Lock()
	defer this.memlock.Unlock()

	// Memory map GPIO registers to byte array
	this.mem8, err = syscall.Mmap(
		int(file.Fd()),
		base,
		memLength,
		syscall.PROT_READ|syscall.PROT_WRITE,
		syscall.MAP_SHARED)

	if err != nil {
		return err
	}

	// Convert mapped byte memory to unsafe []uint32 pointer, adjust length as needed
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&this.mem8))
	header.Len /= (32 / 8) // (32 bit = 4 bytes)
	header.Cap /= (32 / 8)

	this.mem = *(*[]uint32)(unsafe.Pointer(&header))

	return nil
}

func (this *GpioSingleton) Close() error {
	// Clear the singleton variable.
	GpioSingletonInstance = nil
	// Clear all pins.
	this.pins = nil
	// If the Mock flag is set do nothing.
	if this.mock {
		return nil
	}
	// Free the memory.
	this.memlock.Lock()
	defer this.memlock.Unlock()
	return syscall.Munmap(this.mem8)
}

// Returns if this instance is a mock or not.
func (this *GpioSingleton) IsMock() bool {
	return this.mock
}

func (this *GpioSingleton) Pin(pin uint8) *GpioPin {
	if int(pin) > len(this.pins) {
		return nil
	}
	if p := this.pins[pin]; p != nil {
		return p
	}
	this.pins[pin] = NewPin(this, pin)
	return this.pins[pin]
}

func (this *GpioSingleton) Pull(pin uint8, pull Pull) {
	// If the Mock flag is set do nothing.
	if this.mock {
		return
	}
	// Pull up/down/off register has offset 38 / 39, pull is 37
	pullClkReg := uint8(pin)/32 + 38
	pullReg := 37
	shift := (uint8(pin) % 32)

	this.memlock.Lock()
	defer this.memlock.Unlock()

	switch pull {
	case PullDown, PullUp:
		this.mem[pullReg] = this.mem[pullReg]&^3 | uint32(pull)
	case PullOff:
		this.mem[pullReg] = this.mem[pullReg] &^ 3
	}

	// Wait for value to clock in, this is ugly, sorry :(
	time.Sleep(time.Microsecond)

	this.mem[pullClkReg] = 1 << shift

	// Wait for value to clock in
	time.Sleep(time.Microsecond)

	this.mem[pullReg] = this.mem[pullReg] &^ 3
	this.mem[pullClkReg] = 0
}

func (this *GpioSingleton) Mode(pin uint8, dir Direction) {
	// If the Mock flag is set do nothing.
	if this.mock {
		return
	}
	// GpioPin fsel register, 0 or 1 depending on bank
	fsel := uint8(pin) / 10
	shift := (uint8(pin) % 10) * 3

	this.memlock.Lock()
	defer this.memlock.Unlock()

	if dir == Input {
		this.mem[fsel] = this.mem[fsel] &^ (pinMask << shift)
	} else {
		this.mem[fsel] = (this.mem[fsel] &^ (pinMask << shift)) | (1 << shift)
	}
}

func (this *GpioSingleton) Read(pin uint8) State {
	// If the Mock flag is set do nothing.
	if this.mock {
		if GpioPin := this.Pin(pin); GpioPin != nil {
			return GpioPin.LastWrite()
		} else {
			return Low
		}
	}
	// Input level register offset (13 / 14 depending on bank)
	levelReg := uint8(pin)/32 + 13

	if (this.mem[levelReg] & (1 << uint8(pin))) != 0 {
		return High
	}

	return Low
}

func (this *GpioSingleton) Write(pin uint8, state State) {
	// If the Mock flag is set do nothing.
	if this.mock {
		return
	}
	// Clear register, 10 / 11 depending on bank
	// Set register, 7 / 8 depending on bank
	clearReg := pin/32 + 10
	setReg := pin/32 + 7

	this.memlock.Lock()
	defer this.memlock.Unlock()

	if state == Low {
		this.mem[clearReg] = 1 << (pin & 31)
	} else {
		this.mem[setReg] = 1 << (pin & 31)
	}
}

// Read /proc/device-tree/soc/ranges and determine the base address.
// Use the default Raspberry Pi 1 base address if this fails.
func (this *GpioSingleton) getGPIOBase() (base int64) {
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
