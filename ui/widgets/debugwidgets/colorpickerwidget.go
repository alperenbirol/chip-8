package debugwidgets

import (
	"github.com/AllenDang/giu"
	"github.com/alperenbirol/chip-8/emuconfig"
)

func ColorPickerWidget() giu.Layout {
	return giu.Layout{
		giu.Row(
			giu.Column(
				giu.ColorEdit("Primary Color", &emuconfig.PIXEL_ON_COLOR).Flags(giu.ColorEditFlagsHEX|giu.ColorEditFlagsNoLabel).Size(125),
			),
			giu.Spacing(),
			giu.Spacing(),
			giu.Spacing(),
			giu.Spacing(),
			giu.Column(
				giu.ColorEdit("Secondary Color", &emuconfig.PIXEL_OFF_COLOR).Flags(giu.ColorEditFlagsHEX|giu.ColorEditFlagsNoLabel).Size(125),
			),
		),
	}
}
