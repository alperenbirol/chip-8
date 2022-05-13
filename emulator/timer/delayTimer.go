package timer

import (
	"time"

	"github.com/alperenbirol/chip-8/emuconfig"
)

type delayTimer struct {
	timer
}

type IDelayTimer interface {
	SetTimer(d byte)
	Tick()
	GetRemainingTimePointer() *byte
}

func NewDelayTimer() IDelayTimer {
	delayTimer := &delayTimer{
		timer: timer{
			remainingTime: 0x00,
			ticker:        time.NewTicker(emuconfig.TIMER_INTERVAL),
		},
	}

	go delayTimer.Tick()

	return delayTimer
}
