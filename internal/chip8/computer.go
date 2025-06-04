package chip8

import "AlvinSamuelsson/chip8go/internal/font"

const (
	totalMemory = 4096
)

type Chip8 struct {
	Memory [totalMemory]byte
	opcode uint16
	Pc     uint16
}

func NewPC() (*Chip8, error) {
	vm := Chip8{
		Memory: [4096]byte{},
		Pc:     0x200,
	}

	vm.loadFont()

	return &vm, nil
}

func (chip *Chip8) loadFont() {
	// load font into memory, on place 000-1FF
	fontInc := 0
	for i := 0x50; i <= 0x9F; i++ {
		chip.Memory[i] = font.FontSet[fontInc]
		fontInc++
	}
}
