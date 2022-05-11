package timer

type delayTimer struct {
	timer
}

func NewDelayTimer() delayTimer {
	return delayTimer{
		timer: 0x00,
	}
}
