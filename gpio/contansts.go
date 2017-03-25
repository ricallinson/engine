package gpio

type Direction uint8
type State uint8
type Pull uint8

// Memory offsets for gpio, see the spec for more details
const (
	bcm2835Base        = 0x20000000
	pi1GPIOBase        = bcm2835Base + 0x200000
	memLength          = 4096
	pinMask     uint32 = 7 // 0b111 - pinmode is 3 bits
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
