//
// Copyright 2017, Yahoo Inc.
// Copyrights licensed under the New BSD License.
// See the accompanying LICENSE file for terms.
//

package engine

import (
// "github.com/ricallinson/engine/rpio"
)

type Switch struct {
	*PinInput
}

// Returns a new instance of Switch.
// The value of `pin` must be in the range of 1-25 mapping to the Raspberry Pi GPIO pins.
func NewSwitch(pin int) *Switch {
	this := &Switch{
		NewPinInput(pin),
	}
	return this
}
