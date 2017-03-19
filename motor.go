package engine

type Motor struct {
}

func NewMotor(pin int, reversed bool) *Motor {
	this := &Motor{}
	return this
}

func (this *Motor) Set(val float32) {

}
