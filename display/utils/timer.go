package utils

type Timer struct {
	Length int64
	Ring   func(reset func())

	duration int64
}

func (ths *Timer) Tick(delta int64) {
	ths.duration += delta
	if ths.duration > ths.Length {
		ths.Ring(func() { ths.duration = 0 })
	}
}
