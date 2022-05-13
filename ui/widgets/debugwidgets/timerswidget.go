package debugwidgets

import (
	"fmt"

	"github.com/AllenDang/giu"
)

func TimerWidget(timerRemainingTime *byte) *giu.Layout {
	return &giu.Layout{
		giu.Label(fmt.Sprintf("0x%x", *timerRemainingTime)),
	}
}
