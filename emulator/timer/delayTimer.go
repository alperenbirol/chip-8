package timer

type delayTimer struct {
	timer
}

type IDelayTimer interface {
	SetTimer(d byte)
	Tick()
}

func NewDelayTimer() IDelayTimer {
	return &delayTimer{
		timer: 0x00,
	}
}
