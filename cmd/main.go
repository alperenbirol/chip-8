package main

import (
	"github.com/alperenbirol/chip-8/emulator"
	"github.com/alperenbirol/chip-8/emulator/beeper"
	"github.com/alperenbirol/chip-8/ui"
)

func main() {
	beeper, err := beeper.NewBeeper()
	if err != nil {
		panic(err)
	}

	emulator := emulator.NewEmulator(beeper)
	emulator.LoadROM("./roms/ibm.ch8")
	emulator.Run()

	gui := ui.NewGUI(&ui.GuiConfig{
		IsDebugging: true,
		DebugProps:  emulator.DebugProps,
	})
	gui.Run()
}
