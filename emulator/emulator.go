package emulator

import (
	"os"
	"sync"
	"time"

	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/beeper"
	"github.com/alperenbirol/chip-8/emulator/decoder"
	"github.com/alperenbirol/chip-8/emulator/vm"
)

type Emulator struct {
	vm     *vm.VirtualMachine
	ticker *time.Ticker
}

func (e *Emulator) instructionLoop() {
	for range e.ticker.C {
		pc := e.vm.PC.Get()
		instruction := e.vm.RAM.FetchInstruction(pc)

		cmd := decoder.Decode(&instruction)
		cmd(e.vm)

		e.vm.PC.NextInstruction()
	}
}

func (e *Emulator) Run() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		e.instructionLoop()
	}()
	go func() {
		for {
			if e.vm.IsDrawing {
				e.vm.Display.AsciiPrint()
				e.vm.IsDrawing = false
			}
		}
	}()
	wg.Wait()
}

func loadFile(filepath string) ([]byte, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func NewEmulator() *Emulator {
	beeper, err := beeper.NewBeeper()
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(emuconfig.TIMER_INTERVAL)

	return &Emulator{
		vm:     vm.NewVirtualMachine(beeper),
		ticker: ticker,
	}
}

func (e *Emulator) LoadROM(pathToRom string) error {
	bytes, err := loadFile(pathToRom)
	if err != nil {
		return err
	}

	e.vm.RAM.LoadROM(bytes)
	e.vm.PC.SetToAddress(0x200)

	return nil
}

func (e *Emulator) Reset() {
	e.vm.PC.SetToAddress(0x200)
}

func (e *Emulator) Clear() {
	e.vm.Display.Clear()
	e.vm.RAM.UnloadROM()
}

func (e *Emulator) IsDrawing() bool {
	return e.vm.IsDrawing
}

func (e *Emulator) GetDisplayBits() emuconfig.Pixels {
	return e.vm.Display.GetDisplay()
}