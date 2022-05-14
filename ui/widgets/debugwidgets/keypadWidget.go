package debugwidgets

import (
	"github.com/AllenDang/giu"
	"github.com/alperenbirol/chip-8/emulator"
)

func KeypadWindow(functions *emulator.KeypressFunctions) *giu.Layout {
	return &giu.Layout{
		giu.Row(
			giu.Button("1").Size(50, 50).OnClick(functions.KEY_1_PRESSED),
			giu.Button("2").Size(50, 50).OnClick(functions.KEY_2_PRESSED),
			giu.Button("3").Size(50, 50).OnClick(functions.KEY_3_PRESSED),
			giu.Button("C").Size(50, 50).OnClick(functions.KEY_C_PRESSED),
		),
		giu.Row(
			giu.Button("4").Size(50, 50).OnClick(functions.KEY_4_PRESSED),
			giu.Button("5").Size(50, 50).OnClick(functions.KEY_5_PRESSED),
			giu.Button("6").Size(50, 50).OnClick(functions.KEY_6_PRESSED),
			giu.Button("D").Size(50, 50).OnClick(functions.KEY_D_PRESSED),
		),
		giu.Row(
			giu.Button("7").Size(50, 50).OnClick(functions.KEY_7_PRESSED),
			giu.Button("8").Size(50, 50).OnClick(functions.KEY_8_PRESSED),
			giu.Button("9").Size(50, 50).OnClick(functions.KEY_9_PRESSED),
			giu.Button("E").Size(50, 50).OnClick(functions.KEY_E_PRESSED),
		),
		giu.Row(
			giu.Button("A").Size(50, 50).OnClick(functions.KEY_A_PRESSED),
			giu.Button("0").Size(50, 50).OnClick(functions.KEY_0_PRESSED),
			giu.Button("B").Size(50, 50).OnClick(functions.KEY_B_PRESSED),
			giu.Button("F").Size(50, 50).OnClick(functions.KEY_F_PRESSED),
		),
	}
}
