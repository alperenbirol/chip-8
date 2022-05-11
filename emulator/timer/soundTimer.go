package timer

type soundTimer struct {
	timer
}

func NewSoundTimer() soundTimer {
	return soundTimer{
		timer: 0x00,
	}
}

func (t *soundTimer) Tick() {
	t.timer.Tick()

}
