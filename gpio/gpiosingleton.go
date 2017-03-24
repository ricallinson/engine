package gpio

var gpioSingletonInstance *gpioSingleton

type gpioSingleton struct {
	pins []*Pin
}

func Gpio() *gpioSingleton {
	if gpioSingletonInstance != nil {
		return gpioSingletonInstance
	}
	this := &gpioSingleton{}
	gpioSingletonInstance = this
	return this
}

func (this *gpioSingleton) Open() {

}

func (this *gpioSingleton) Close() {

}

func (this *gpioSingleton) Pin(pin uint8) *Pin {
	if pin := this.pins[pin]; pin != nil {
		return pin
	}
	this.pins[pin] = NewPin(this, pin)
	return this.pins[pin]
}
