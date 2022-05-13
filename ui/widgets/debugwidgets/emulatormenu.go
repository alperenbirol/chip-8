package debugwidgets

import (
	"github.com/AllenDang/giu"
	"github.com/alperenbirol/chip-8/emulator"
)

func EmulatorMenuWidget(functions *emulator.EmulatorFunctions, paused bool, freq *int32) giu.Layout {
	var pauseResumeButton *giu.ButtonWidget
	if paused {
		pauseResumeButton = giu.Button("Resume").OnClick(functions.Resume)
	} else {
		pauseResumeButton = giu.Button("Pause").OnClick(functions.Pause)
	}

	return giu.Layout{
		giu.Spacing(),
		giu.Spacing(),
		giu.Row(
			giu.Column(
				giu.Button("Reset").Size(150, 50).OnClick(functions.Reset),
			),
			giu.SliderInt(freq, 2, 300).Label("Clock Speed").OnChange(func() {
				functions.SetClockSpeed(uint(*freq))
			}),
		),
		giu.Row(
			giu.Column(
				pauseResumeButton.Size(150, 50),
			),
		),
		giu.Row(
			giu.Column(
				giu.Button("Step").Size(150, 50).Disabled(!paused).OnClick(functions.Step),
			),
		),
	}
}
