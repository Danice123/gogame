package battle

type BattlesystemListener interface {
	Text(text string)
	WaitForInput()
}

type Battlesystem struct {
	Listener BattlesystemListener
}

func (ths *Battlesystem) Start() {

}
