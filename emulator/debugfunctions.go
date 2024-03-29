package emulator

import (
	"time"

	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/indexregister"
	"github.com/alperenbirol/chip-8/emulator/programregister"
	"github.com/alperenbirol/chip-8/emulator/stack"
)

type EmulatorFunctions struct {
	Reset         func()
	LoadROM       func(string) error
	Clear         func()
	Pause         func()
	Resume        func()
	SetClockSpeed func(uint)
	Step          func()
	SetKey        func(emuconfig.Key, bool)
	GetDisplay    func() emuconfig.Pixels
}

type DebugProps struct {
	Instructions    []emuconfig.Opcode
	Memory          *emuconfig.Ram
	Registers       *[16]programregister.ProgramRegister
	IndexRegister   *indexregister.IndexRegister
	Stack           *stack.Stack
	ProgramCounter  *uint16
	DelayTimerTime  *byte
	SoundTimerTime  *byte
	IsDrawing       *bool
	Paused          *bool
	Functions       *EmulatorFunctions
	KeypadFunctions *KeypressFunctions
}

func (e *Emulator) setDebugProps() {
	e.DebugProps.Paused = &e.paused
	e.DebugProps.Memory = e.vm.RAM.GetRamPointer()
	e.DebugProps.Registers = &e.vm.Registers
	e.DebugProps.IndexRegister = &e.vm.IndexRegister
	e.DebugProps.Stack = e.vm.Stack
	e.DebugProps.IsDrawing = &e.vm.IsDrawing
	e.DebugProps.ProgramCounter = e.vm.PC.GetPointer()
	e.DebugProps.DelayTimerTime = e.vm.DelayTimer.GetRemainingTimePointer()
	e.DebugProps.SoundTimerTime = e.vm.SoundTimer.GetRemainingTimePointer()
}

func (e *Emulator) setDebugFunctions() {
	e.DebugProps.Functions = &EmulatorFunctions{
		Reset:         e.Reset,
		LoadROM:       e.LoadROM,
		Clear:         e.Clear,
		GetDisplay:    e.GetDisplayBits,
		Pause:         e.Pause,
		Resume:        e.Resume,
		Step:          e.instructionStep,
		SetClockSpeed: e.SetClockSpeed,
		SetKey:        e.SetKey,
	}
	e.DebugProps.KeypadFunctions = e.KeypressFunctions()
}

func (e *Emulator) Reset() {
	e.vm.IndexRegister.Set(0x00)
	e.vm.Registers = [16]programregister.ProgramRegister{}
	e.vm.DelayTimer.SetTimer(0x00)
	e.vm.SoundTimer.SetTimer(0x00)
	e.vm.PC.SetToAddress(0x200)
	e.DebugProps.Instructions = []emuconfig.Opcode{}
}

func (e *Emulator) Clear() {
	e.vm.Display.Clear()
	e.vm.RAM.UnloadROM()
	e.Reset()
	e.paused = true
}

func (e *Emulator) GetDisplayBits() emuconfig.Pixels {
	return e.vm.Display.GetDisplay()
}

func (e *Emulator) Pause() {
	e.paused = true
}

func (e *Emulator) Resume() {
	e.paused = false
}

func (e *Emulator) SetKey(key emuconfig.Key, pressed bool) {
	if pressed {
		e.vm.Keypad.PressKey(key.KeyCode)
	} else {
		e.vm.Keypad.ReleaseKey(key.KeyCode)
	}
}

func (e *Emulator) SetClockSpeed(frequency uint) {
	e.ticker.Reset(time.Second / time.Duration(frequency))
}

func (e *Emulator) addInstruction(instruction emuconfig.Opcode) {
	{
		if len(e.DebugProps.Instructions) == 0 || e.DebugProps.Instructions[len(e.DebugProps.Instructions)-1] != instruction {
			e.DebugProps.Instructions = append(e.DebugProps.Instructions, instruction)
		}
	}
}
