package chip8

import (
	"AlvinSamuelsson/chip8go/internal/font"
	"AlvinSamuelsson/chip8go/internal/gui"
	"fmt"
	"os"
	"time"
)

const (
	totalMemory = 4096
	vRegister   = 16
)

type Chip8 struct {
	Memory    [totalMemory]byte
	Gfx       [64 * 32]byte
	Vregister [vRegister]uint8
	Iregister uint16
	// Stack  Stack
	opcode uint16
	Pc     uint16
	dTimer byte
	sTimer byte
	Clock  *time.Ticker
	window *gui.Window
}

func NewPC(path string) (*Chip8, error) {
	window, err := gui.CreateWindow()

	if err != nil {
		panic("Creating window failed!")
	}

	vm := Chip8{
		Memory: [4096]byte{},
		Pc:     0x200,
		dTimer: 0xFF,
		sTimer: 0xFF,
		Clock:  time.NewTicker(time.Second / time.Duration(60)),
		window: window,
	}

	vm.loadFont()
	vm.loadProgram(path)
	return &vm, nil
}

func (chip *Chip8) RunProgram() {
outer:
	for {
		select {
		case <-chip.Clock.C:
			if !chip.window.Closed() {
				chip.cycle()
				chip.DrawGraphics()
				//chip.HandleInput()
				chip.delayTimer()
				chip.soundTimer()
				continue
			}
			break outer
		}
	}
}

func (chip *Chip8) loadProgram(path string) error {
	rom, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	//load program
	for i := range len(rom) {
		chip.Memory[0x200+i] = rom[i]
	}

	return nil
}

func (chip *Chip8) cycle() {
	chip.opcode = uint16(chip.Memory[chip.Pc])<<8 | uint16(chip.Memory[chip.Pc+1])
	var x = (chip.opcode & 0x0F00) >> 8
	// var y = (chip.opcode & 0x00F0) >> 4
	var nn uint8 = uint8(chip.opcode & 0x00FF)
	var nnn uint16 = (chip.opcode & 0x0FFF)
	fmt.Printf("Cycle PC=%x --- Opcode =0x%04x --- instruction =0x%04x \n", chip.Pc, chip.opcode, (chip.opcode & 0xF000))

	chip.Pc = chip.Pc + 2

	switch (chip.opcode & 0xf000) << 12 {
	case 0x0:
		//clear screencomp
		switch chip.opcode {
		case 0x0E0:
			fmt.Println("We are clearing da screen!")
		}
	case 0x100:
		chip._0x100(nnn)
	case 0x600:
		chip.Vregister[x] = nn
	case 0x700:
		chip.Vregister[x] += nn
	case 0xa00:
		chip.Iregister = nnn
	}
}

func (chip *Chip8) DrawGraphics() {
	chip.window.DrawGraphics(chip.Gfx)
}

func (chip *Chip8) loadFont() error {
	// load font into memory, on place 000-1FF
	fontInc := 0
	for i := 0x50; i <= 0x9F; i++ {
		chip.Memory[i] = font.FontSet[fontInc]
		fontInc++
	}

	return nil
}

func (chip *Chip8) soundTimer() {
	if chip.sTimer != 0 {
		chip.sTimer--
	}
}

func (chip *Chip8) delayTimer() {
	if chip.dTimer != 0 {
		chip.dTimer--
	}
}
