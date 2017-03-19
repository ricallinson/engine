package engine

type IRSensor struct {
}

func NewIRSensor(pin int) *IRSensor {
	this := &IRSensor{}
	return this
}

func (this *IRSensor) Get() float32 {
	return 0
}
