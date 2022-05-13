package main

import (
	"image/color"

	g "github.com/AllenDang/giu"
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator"
	"github.com/alperenbirol/chip-8/emulator/programregister"
	"github.com/alperenbirol/chip-8/ui/displayconverter"
	"github.com/alperenbirol/chip-8/ui/widgets"
	"github.com/alperenbirol/chip-8/ui/widgets/debugwidgets"
	"github.com/alperenbirol/chip-8/ui/windows/mainwindow"
)

var gui *GUI
var ran = false

type debug struct {
	instructions []emuconfig.Opcode
	memory       emuconfig.Ram
	registers    [16]programregister.ProgramRegister
}

type GUI struct {
	display  *g.Texture
	emulator *emulator.Emulator

	isDebugging bool
	debug
}

func loop() {
	if !ran {
		setTextureFilter()
	}
	go gui.refreshDisplay()
	if gui.isDebugging {
		go gui.getDebugProps()
		g.Window("Registers").Size(1650, 75).Pos(0, 0).Flags(g.WindowFlagsNoCollapse | g.WindowFlagsNoInputs).Layout(
			debugwidgets.RegistersWidget(gui.debug.registers),
		)
		g.Window("Memory").Size(435, 360).Pos(0, 75).Flags(g.WindowFlagsNoCollapse | g.WindowFlagsNoInputs).Layout(
			debugwidgets.MemoryWidget(gui.debug.memory),
		)
		g.Window("Keypad").Size(250, 250).Pos(1090, 75).Flags(g.WindowFlagsNoCollapse | g.WindowFlagsNoInputs).Layout(
			debugwidgets.KeypadWindow(),
		)
		g.Window("Instructions").Size(650, 310).Pos(1090, 325).Flags(g.WindowFlagsNoCollapse | g.WindowFlagsNoInputs).Layout(
			debugwidgets.InstructionsWidget(gui.instructions),
		)
	}

	g.Window("Display").Pos(435, 75).Size(655, 360).Flags(g.WindowFlagsNoCollapse | g.WindowFlagsNoResize | g.WindowFlagsNoMove).Layout(
		widgets.DisplayWidget(gui.display).Size(640, 320).BorderCol(color.White),
	)
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

func (gui *GUI) getDebugProps() {
	debug := gui.emulator.DebugProps
	for {
		lastInstruction := <-debug.Instruction
		if len(gui.debug.instructions) == 0 || gui.debug.instructions[len(gui.debug.instructions)-1] != lastInstruction {
			gui.debug.instructions = append(gui.debug.instructions, lastInstruction)
		}
		gui.debug.memory = debug.Memory
		gui.debug.registers = debug.Registers
	}
}

func setTextureFilter() error {
	err := g.Context.GetRenderer().SetTextureMagFilter(g.TextureFilterNearest)
	if err != nil {
		return err
	}

	return nil
}
