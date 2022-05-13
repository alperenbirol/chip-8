package debugwidgets

import (
	"github.com/AllenDang/giu"
	"github.com/alperenbirol/chip-8/emulator"
)

func EmulatorMenuWidget(functions *emulator.EmulatorFunctions) giu.Layout {
	return giu.Layout{
		giu.Button("Reset").Size(150, 50).OnClick(functions.Reset),
	}
}
