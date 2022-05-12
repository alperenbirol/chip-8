package main

import (
	"image/color"

	g "github.com/AllenDang/giu"
	"github.com/alperenbirol/chip-8/emulator"
	"github.com/alperenbirol/chip-8/ui/displayconverter"
	"github.com/alperenbirol/chip-8/ui/widgets"
	"github.com/alperenbirol/chip-8/ui/widgets/debugwidgets"
	"github.com/alperenbirol/chip-8/ui/windows/mainwindow"
)

var gui *GUI
var ran = false

type GUI struct {
	display     *g.Texture
	emulator    *emulator.Emulator
	isDebugging bool
}

func loop() {
	if !ran {
		setTextureFilter()
	}
	go gui.refreshDisplay()
	if gui.isDebugging {
		g.Window("Registers").Size(1650, 75).Pos(0, 0).Flags(g.WindowFlagsNoCollapse | g.WindowFlagsNoInputs).Layout(
			debugwidgets.RegistersWidget(gui.emulator.GetRegisters()),
		)
		g.Window("Memory").Size(435, 240).Pos(0, 75).Flags(g.WindowFlagsNoCollapse | g.WindowFlagsNoInputs).Layout(
			debugwidgets.MemoryWidget(gui.emulator.GetMemory()),
		)
	}

	g.Window("Display").Pos(435, 75).Size(345, 200).Flags(g.WindowFlagsNoCollapse | g.WindowFlagsNoResize).Layout(
		widgets.DisplayWidget(gui.display).Size(320, 160).BorderCol(color.White),
	)
	// if in debug mode
}

func main() {
	gui = &GUI{
		emulator:    emulator.NewEmulator(),
		isDebugging: true,
	}

	gui.emulator.LoadROM("../roms/ibm.ch8")
	gui.emulator.Run()

	mainwindow.InitGUI(loop)
}

func (gui *GUI) refreshDisplay() {
	if gui.emulator.IsDrawing() {
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
