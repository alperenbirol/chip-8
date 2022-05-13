package main

import (
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator"
	"github.com/alperenbirol/chip-8/emulator/beeper"
	"github.com/alperenbirol/chip-8/ui"
	"github.com/alperenbirol/chip-8/ui/widgets/debugwidgets"
)

func main() {
	beeper, err := beeper.NewBeeper()
	if err != nil {
		panic(err)
	}

	romPath := "./roms/ibm.ch8"

	emulator := emulator.NewEmulator(beeper)
	emulator.LoadROM(romPath)
	emulator.Run()

	var freq int32 = emuconfig.TIMER_FREQUENCY
	gui := ui.NewGUI(&ui.GuiConfig{
		IsDebugging: true,
		DebugProps:  emulator.DebugProps,
		EmulatorDebugMenuProps: &debugwidgets.EmulatorMenuProps{
			Paused:  emulator.DebugProps.Paused,
			Freq:    &freq,
			RomPath: &romPath,
		},
	})
	gui.Run()
}
