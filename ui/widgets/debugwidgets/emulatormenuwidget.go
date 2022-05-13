package debugwidgets

import (
	"github.com/AllenDang/giu"
	"github.com/alperenbirol/chip-8/emulator"
	"github.com/alperenbirol/chip-8/ui/filebrowser"
	"github.com/sqweek/dialog"
)

type EmulatorMenuProps struct {
	Paused  *bool
	Freq    *int32
	RomPath *string
}

func EmulatorMenuWidget(functions *emulator.EmulatorFunctions, props *EmulatorMenuProps) giu.Layout {
	var pauseResumeButton *giu.ButtonWidget
	if *props.Paused {
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
			giu.Column(
				giu.Label("Clock Speed"),
			),
			giu.Column(
				giu.SliderInt(props.Freq, 1, 300).Size(230).OnChange(func() {
					functions.SetClockSpeed(uint(*props.Freq))
				}),
			),
		),
		giu.Row(
			giu.Column(
				pauseResumeButton.Size(150, 50),
			),
			giu.Column(
				giu.Row(
					giu.Label("ROM Path"),
					giu.InputText(props.RomPath).Size(140),
					giu.Button("Open").Size(50, 20).OnClick(func() {
						romPath, err := filebrowser.BrowseRom()
						if err != nil {
							if err == dialog.ErrCancelled {
								return
							}
							panic(err)
						}
						props.RomPath = &romPath
					}),
					giu.Button("Load").Size(50, 20).OnClick(func() {
						functions.LoadROM(*props.RomPath)
					}),
				),
			),
		),
		giu.Row(
			giu.Column(
				giu.Button("Step").Size(150, 50).Disabled(!*props.Paused).OnClick(functions.Step),
			),
		),
	}
}
