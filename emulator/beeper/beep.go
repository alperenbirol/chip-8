package beeper

import (
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/faiface/beep"
	"github.com/faiface/beep/generators"
	"github.com/faiface/beep/speaker"
)

type beeper struct {
	tone       *beep.Streamer
	beeperCtrl *beep.Ctrl
}

func NewBeeper() (*beeper, error) {
	err := speaker.Init(emuconfig.BEEPER_SAMPLE_RATE, emuconfig.BEEPER_BUFFER_SIZE)
	if err != nil {
		return nil, err
	}

	sine, err := generators.SinTone(emuconfig.BEEPER_SAMPLE_RATE, emuconfig.BEEPER_FREQUENCY)
	if err != nil {
		return nil, err
	}

	ctrl := &beep.Ctrl{Streamer: sine, Paused: true}

	speaker.Play(ctrl)

	return &beeper{
		tone:       &sine,
		beeperCtrl: ctrl,
	}, nil
}

func (b *beeper) Play() {
	b.beeperCtrl.Paused = false
}

func (b *beeper) Stop() {
	b.beeperCtrl.Paused = true
}
