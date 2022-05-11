package timer

import (
	"time"

	"github.com/alperenbirol/chip-8/emuconfig"
)

type timer byte

func (t *timer) SetTimer(d byte) {
	*t = timer(d)
}

func (t *timer) Tick() {
	go func(t *timer) {
		for {
			if *t > 0 {
				t.decrease()
			}
		}
	}(t)
}

func (t *timer) decrease() {
	*t--
	time.Sleep(time.Second / emuconfig.TIMER_FREQUENCY)
}
