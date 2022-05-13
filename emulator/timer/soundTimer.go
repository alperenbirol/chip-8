package timer

import (
	"time"

	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/beeper"
)

type soundTimer struct {
	timer
	beeper.IBeeper
}

type ISoundTimer interface {
	SetTimer(d byte)
	Tick()
}

func NewSoundTimer(beeper beeper.IBeeper) ISoundTimer {
	soundTimer := &soundTimer{
		timer: timer{
			remainingTime: 0x00,
			ticker:        time.NewTicker(emuconfig.TIMER_INTERVAL),
		},
		IBeeper: beeper,
	}

	go soundTimer.Tick()

	return soundTimer
}

func (t *soundTimer) SetTimer(d byte) {
	t.remainingTime = d
	if d > 0 {
		t.Play()
	}
}

func (t *soundTimer) Tick() {
	for {
		<-t.ticker.C
		if t.remainingTime > 0 {
			t.decrease()
		}
		if t.remainingTime == 0 {
			t.Stop()
		}
	}
}
