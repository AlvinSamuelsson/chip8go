package chip8

import (
	"AlvinSamuelsson/chip8go/internal/font"
	"os"
)

const (
	totalMemory = 4096
)

type Chip8 struct {
	Memory [totalMemory]byte
	// Stack  Stack
	opcode uint16
	Pc     uint16
}

func NewPC(path string) (*Chip8, error) {
	vm := Chip8{
		Memory: [4096]byte{},
		Pc:     0x200,
	}

	vm.loadFont()
	vm.loadProgram(path)
	vm.cycle()

	return &vm, nil
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
	chip.opcode = uint16(chip.Memory[chip.Pc]<<8) | uint16(chip.Memory[chip.Pc+1])
	var x = (chip.opcode & 0x0F00) >> 8
	var y = (chip.opcode & 0x00F0) >> 4
	var nn = (chip.opcode & 0x00FF)
	var nnn = (chip.opcode & 0x0FFF)

	switch (chip.opcode & 0xF000) << 12 {
	case 0x00E0:
		//clear screen
	case 0x100:
		chip.Pc = nnn
	}
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
