package debugwidgets

import (
	"github.com/AllenDang/giu"
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator"
)

func KeypadWindow(functions *emulator.KeypressFunctions) *giu.Layout {
	return &giu.Layout{
		giu.Row(
			giu.Button("1").Size(50, 50).OnClick(functions.KEY_1_PRESSED).Disabled(isKeyPressed(emuconfig.KEY_1)),
			giu.Button("2").Size(50, 50).OnClick(functions.KEY_2_PRESSED).Disabled(isKeyPressed(emuconfig.KEY_2)),
			giu.Button("3").Size(50, 50).OnClick(functions.KEY_3_PRESSED).Disabled(isKeyPressed(emuconfig.KEY_3)),
			giu.Button("C").Size(50, 50).OnClick(functions.KEY_C_PRESSED).Disabled(isKeyPressed(emuconfig.KEY_C)),
		),
		giu.Row(
			giu.Button("4").Size(50, 50).OnClick(functions.KEY_4_PRESSED).Disabled(isKeyPressed(emuconfig.KEY_4)),
			giu.Button("5").Size(50, 50).OnClick(functions.KEY_5_PRESSED).Disabled(isKeyPressed(emuconfig.KEY_5)),
			giu.Button("6").Size(50, 50).OnClick(functions.KEY_6_PRESSED).Disabled(isKeyPressed(emuconfig.KEY_6)),
			giu.Button("D").Size(50, 50).OnClick(functions.KEY_D_PRESSED).Disabled(isKeyPressed(emuconfig.KEY_D)),
		),
		giu.Row(
			giu.Button("7").Size(50, 50).OnClick(functions.KEY_7_PRESSED).Disabled(isKeyPressed(emuconfig.KEY_7)),
			giu.Button("8").Size(50, 50).OnClick(functions.KEY_8_PRESSED).Disabled(isKeyPressed(emuconfig.KEY_8)),
			giu.Button("9").Size(50, 50).OnClick(functions.KEY_9_PRESSED).Disabled(isKeyPressed(emuconfig.KEY_9)),
			giu.Button("E").Size(50, 50).OnClick(functions.KEY_E_PRESSED).Disabled(isKeyPressed(emuconfig.KEY_E)),
		),
		giu.Row(
			giu.Button("A").Size(50, 50).OnClick(functions.KEY_A_PRESSED).Disabled(isKeyPressed(emuconfig.KEY_A)),
			giu.Button("0").Size(50, 50).OnClick(functions.KEY_0_PRESSED).Disabled(isKeyPressed(emuconfig.KEY_0)),
			giu.Button("B").Size(50, 50).OnClick(functions.KEY_B_PRESSED).Disabled(isKeyPressed(emuconfig.KEY_B)),
			giu.Button("F").Size(50, 50).OnClick(functions.KEY_F_PRESSED).Disabled(isKeyPressed(emuconfig.KEY_F)),
		),
	}
}

func isKeyPressed(key emuconfig.Key) bool {
	return giu.IsKeyPressed(emuconfig.REVERSE_USER_KEYPAD_INTERFACE[key])
}
