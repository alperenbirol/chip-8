package debugwidgets

import "github.com/AllenDang/giu"

func EmulatorMenuWidget(reset func()) giu.Layout {
	return giu.Layout{
		giu.Button("Reset").Size(150, 50).OnClick(reset),
	}
}
