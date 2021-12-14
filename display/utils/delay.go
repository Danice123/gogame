package utils

type delayobject struct {
	timer  int
	delay  int
	action func()
}

type DelayHandler struct {
	tracked []delayobject
}

func NewDelayHandler() *DelayHandler {
	return &DelayHandler{
		tracked: []delayobject{},
	}
}

func (ths *DelayHandler) Tick() {
	unfinished := []delayobject{}
	for _, ob := range ths.tracked {
		ob.timer++
		if ob.timer >= ob.delay {
			ob.action()
		} else {
			unfinished = append(unfinished, ob)
		}
	}
	ths.tracked = unfinished
}

func (ths *DelayHandler) AddDelayedAction(delay int, action func()) {
	ths.tracked = append(ths.tracked, delayobject{
		delay:  delay,
		action: action,
	})
}
