package emulator

import (
	"os"
	"time"

	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/beeper"
	"github.com/alperenbirol/chip-8/emulator/decoder"
	"github.com/alperenbirol/chip-8/emulator/vm"
)

type Emulator struct {
	vm         *vm.VirtualMachine
	ticker     *time.Ticker
	paused     bool
	DebugProps *DebugProps
}

func (e *Emulator) instructionStep() {
	pc := e.vm.PC.Get()
	instruction := e.vm.RAM.FetchInstruction(pc)
	e.addInstruction(instruction)

	cmd := decoder.Decode(&instruction)
	cmd(e.vm)

	e.vm.PC.NextInstruction()
}

func (e *Emulator) instructionLoop() {
	for range e.ticker.C {
		if !e.paused {
			e.instructionStep()
		}
	}
}

func (e *Emulator) Run() {
	e.paused = false
	go e.instructionLoop()
}

func loadFile(filepath string) ([]byte, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func NewEmulator(beeper beeper.IBeeper) *Emulator {
	ticker := time.NewTicker(emuconfig.TIMER_INTERVAL)

	emulator := &Emulator{
		vm:         vm.NewVirtualMachine(beeper),
		ticker:     ticker,
		paused:     true,
		DebugProps: &DebugProps{},
	}

	emulator.vm.PC.SetToAddress(0x200)
	emulator.vm.RAM.LoadFonts(emuconfig.DEFAULT_FONT_SET[:])

	emulator.setDebugProps()
	emulator.setDebugFunctions()

	return emulator
}

func (e *Emulator) LoadROM(pathToRom string) error {
	e.Clear()

	bytes, err := loadFile(pathToRom)
	if err != nil {
		return err
	}

	e.vm.RAM.LoadROM(bytes)
	e.vm.PC.SetToAddress(0x200)

	return nil
}

func (e *Emulator) PrintBlocking() {
	for range e.ticker.C {
		if e.vm.IsDrawing {
			e.vm.Display.AsciiPrint()
			e.vm.IsDrawing = false
		}
	}
}
