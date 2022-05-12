package main

import (
	"fmt"

	"github.com/alperenbirol/chip-8/emulator"
)

func main() {
	emulator := emulator.NewEmulator()

	emulator.LoadROM("./roms/ibm.ch8")
	emulator.Run()

	for {
		if emulator.IsDrawing() {
			fmt.Println(emulator.GetDisplayBits())
		}
	}

}
