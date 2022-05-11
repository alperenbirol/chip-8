package timer

import "github.com/alperenbirol/chip-8/emulator/beeper"

type soundTimer struct {
	timer
	beeper.IBeeper
}

func NewSoundTimer(beeper beeper.IBeeper) soundTimer {
	return soundTimer{
		timer:   0x00,
		IBeeper: beeper,
	}
}

func (t *soundTimer) Tick() {
	go func(t *soundTimer) {
		for {
			if t.timer > 0 {
				t.decrease()
				t.Play()
			} else {
				t.Stop()
			}
		}
	}(t)
}
