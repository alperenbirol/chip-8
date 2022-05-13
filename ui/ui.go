package main

import (
	"image/color"

	g "github.com/AllenDang/giu"
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator"
	"github.com/alperenbirol/chip-8/emulator/beeper"
	"github.com/alperenbirol/chip-8/ui/displayconverter"
	"github.com/alperenbirol/chip-8/ui/widgets"
	"github.com/alperenbirol/chip-8/ui/widgets/debugwidgets"
	"github.com/alperenbirol/chip-8/ui/windows/mainwindow"
)

var gui *GUI
var ran = false

type GUI struct {
	display  *g.Texture
	emulator *emulator.Emulator

	isDebugging bool
}

func loop() {
	if !ran {
		setTextureFilter()
	}
	go gui.refreshDisplay()
	if gui.isDebugging {
		g.Window("Registers").Size(1650, 80).Pos(0, 0).Flags(emuconfig.DEBUG_WIDGET_FLAGS).Layout(
			debugwidgets.RegistersWidget(gui.emulator.DebugProps.Registers),
		)
		g.Window("Memory").Size(435, 360).Pos(0, 80).Flags(emuconfig.DEBUG_WIDGET_FLAGS).Layout(
			debugwidgets.MemoryWidget(gui.emulator.DebugProps.Memory),
		)
		g.Window("Keypad").Size(250, 250).Pos(1090, 80).Flags(emuconfig.DEBUG_WIDGET_FLAGS).Layout(
			debugwidgets.KeypadWindow(),
		)
		g.Window("Instructions").Size(560, 320).Pos(1090, 330).Flags(emuconfig.DEBUG_WIDGET_FLAGS).Layout(
			debugwidgets.InstructionsWidget(gui.emulator.DebugProps.Instructions),
		)
		g.Window("Index Register").Size(110, 50).Pos(1540, 80).Flags(emuconfig.DEBUG_WIDGET_FLAGS).Layout(
			debugwidgets.IndexRegisterWidget(gui.emulator.DebugProps.IndexRegister),
		)
		g.Window("Emulator Menu").Size(500, 210).Pos(590, 440).Flags(emuconfig.DEBUG_WIDGET_FLAGS).Layout(
			debugwidgets.EmulatorMenuWidget(gui.emulator.Reset),
		)
	}

	g.Window("Display").Pos(435, 80).Size(655, 360).Flags(g.WindowFlagsNoCollapse | g.WindowFlagsNoResize | g.WindowFlagsNoMove).Layout(
		widgets.DisplayWidget(gui.display).Size(640, 320).BorderCol(color.White),
	)
}

func main() {
	beeper, _ := beeper.NewBeeper()
	gui = &GUI{
		emulator:    emulator.NewEmulator(beeper),
		isDebugging: true,
	}

	gui.emulator.LoadROM("../roms/ibm.ch8")
	gui.emulator.Run()

	mainwindow.InitGUI(loop)
}

func (gui *GUI) refreshDisplay() {
	if gui.emulator.DebugProps.IsDrawing {
		g.NewTextureFromRgba(displayconverter.Convert(gui.emulator.GetDisplayBits()), func(t *g.Texture) {
			gui.display = t
		})
	}
	ran = true
}

func setTextureFilter() error {
	err := g.Context.GetRenderer().SetTextureMagFilter(g.TextureFilterNearest)
	if err != nil {
		return err
	}

	return nil
}
