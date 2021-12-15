package utils

type delayobject struct {
	timer  int
	delay  int
	action func()
}

type DelayHandler struct {
	tracked map[*delayobject]bool
}

func NewDelayHandler() *DelayHandler {
	return &DelayHandler{
		tracked: map[*delayobject]bool{},
	}
}

func (ths *DelayHandler) Tick() {
	for ob := range ths.tracked {
		ob.timer++
		if ob.timer >= ob.delay {
			ob.action()
			delete(ths.tracked, ob)
		}
	}
}

func (ths *DelayHandler) AddDelayedAction(delay int, action func()) {
	ob := &delayobject{
		delay:  delay,
		action: action,
	}
	ths.tracked[ob] = true
}
