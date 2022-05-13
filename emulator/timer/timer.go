package timer

import (
	"time"
)

type timer struct {
	remainingTime byte
	ticker        *time.Ticker
}

func (t *timer) SetTimer(d byte) {
	t.remainingTime = d
}

func (t *timer) Tick() {
	for {
		<-t.ticker.C
		if t.remainingTime > 0 {
			t.decrease()
		}
	}
}

func (t *timer) decrease() {
	t.remainingTime--
}

func (t *timer) GetRemainingTimePointer() *byte {
	return &t.remainingTime
}
