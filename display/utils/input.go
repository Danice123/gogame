package utils

var inputList = []KEY{
	ACTIVATE,
	DECLINE,
	UP,
	DOWN,
	LEFT,
	RIGHT,
}

type InputHandler struct {
	inputDelay uint64
	inputGate  uint64
	state      map[KEY]bool

	PressHandlers   map[KEY]func() uint64
	ReleaseHandlers map[KEY]func() uint64
}

func (ths *InputHandler) Tick() {
	ths.inputDelay++
}

func (ths *InputHandler) HandleKey(pressed func(KEY) bool) {
	if ths.state == nil {
		ths.state = map[KEY]bool{}
	}

	ignoreInput := ths.inputDelay < ths.inputGate

	for _, key := range inputList {
		if pressed(key) {
			if !ths.state[key] { // Just Pressed
				ths.state[key] = true
			}

			if handler, ok := ths.PressHandlers[key]; ok && !ignoreInput {
				ths.inputGate = handler()
				ths.inputDelay = 0
				if ths.inputGate > 0 {
					ignoreInput = true
				}
			}

		} else if !pressed(key) && ths.state[key] {
			ths.state[key] = false
			if handler, ok := ths.ReleaseHandlers[key]; ok {
				ths.inputGate = handler()
				ths.inputDelay = 0
				if ths.inputGate > 0 {
					ignoreInput = true
				}
			}
		}
	}
}
