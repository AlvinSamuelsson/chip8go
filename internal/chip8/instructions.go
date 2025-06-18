package chip8

func (chip *Chip8) _0x100(nnn uint16) {
	chip.Pc = nnn
}

func (chip *Chip8) _0x600(x uint16, nn uint8) {
	chip.Vregister[x] = nn
}

func (chip *Chip8) _0x700(x uint16, nn uint8) {
	chip.Vregister[x] += nn
}

func (chip *Chip8) _0xa00(nnn uint16) {
	chip.Iregister = nnn
}
