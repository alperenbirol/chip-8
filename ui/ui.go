package ui

import (
	"image/color"

	g "github.com/AllenDang/giu"
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator"
	"github.com/alperenbirol/chip-8/ui/displayconverter"
	"github.com/alperenbirol/chip-8/ui/widgets"
	"github.com/alperenbirol/chip-8/ui/widgets/debugwidgets"
	"github.com/alperenbirol/chip-8/ui/windows/mainwindow"
)

var ran = false

type GuiConfig struct {
	IsDebugging bool
	DebugProps  *emulator.DebugProps
}

type GUI struct {
	display    *g.Texture
	debugProps *emulator.DebugProps

	isDebugging bool
}

func NewGUI(config *GuiConfig) *GUI {
	return &GUI{
		debugProps:  config.DebugProps,
		isDebugging: config.IsDebugging,
	}
}

func (gui *GUI) loop() {
	if !ran {
		setTextureFilter()
	}
	go gui.refreshDisplay()
	if gui.isDebugging {
		g.Window("Registers").Size(1650, 80).Pos(0, 0).Flags(emuconfig.DEBUG_WIDGET_FLAGS).Layout(
			debugwidgets.RegistersWidget(gui.debugProps.Registers),
		)
		g.Window("Memory").Size(435, 360).Pos(0, 80).Flags(emuconfig.DEBUG_WIDGET_FLAGS).Layout(
			debugwidgets.MemoryWidget(gui.debugProps.Memory),
		)
		g.Window("Keypad").Size(250, 250).Pos(1090, 80).Flags(emuconfig.DEBUG_WIDGET_FLAGS).Layout(
			debugwidgets.KeypadWindow(),
		)
		g.Window("Instructions").Size(560, 320).Pos(1090, 330).Flags(emuconfig.DEBUG_WIDGET_FLAGS).Layout(
			debugwidgets.InstructionsWidget(gui.debugProps.Instructions),
		)
		g.Window("Index Register").Size(110, 50).Pos(1540, 80).Flags(emuconfig.DEBUG_WIDGET_FLAGS).Layout(
			debugwidgets.IndexRegisterWidget(gui.debugProps.IndexRegister),
		)
		g.Window("Emulator Menu").Size(500, 210).Pos(590, 440).Flags(emuconfig.DEBUG_WIDGET_FLAGS).Layout(
			debugwidgets.EmulatorMenuWidget(gui.debugProps.Functions),
		)
	}

	g.Window("Display").Pos(435, 80).Size(655, 360).Flags(g.WindowFlagsNoCollapse | g.WindowFlagsNoResize | g.WindowFlagsNoMove).Layout(
		widgets.DisplayWidget(gui.display).Size(640, 320).BorderCol(color.White),
	)
}

func (gui *GUI) refreshDisplay() {
	if gui.debugProps.IsDrawing {
		g.NewTextureFromRgba(displayconverter.Convert(gui.debugProps.Functions.GetDisplay()), func(t *g.Texture) {
			gui.display = t
		})
	}
}

func setTextureFilter() error {
	err := g.Context.GetRenderer().SetTextureMagFilter(g.TextureFilterNearest)
	if err != nil {
		return err
	}
	ran = true

	return nil
}

func (gui *GUI) Run() {
	mainwindow.InitGUI(gui.loop)
	g.EnqueueNewTextureFromRgba(displayconverter.Convert(gui.debugProps.Functions.GetDisplay()), func(t *g.Texture) {
		gui.display = t
	})
}
