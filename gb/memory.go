package gb

type Memory struct {
	Cart  [0x8000]Word
	VRAM  [0x2000]Word
	ExRAM [0x2000]Word
	WRAM  [0x1000]Word
	// -- prohibited --
	OAM [0xA0]Word
	// -- unusable --
	IO   [0x80]Word
	HRAM [0x7E]Word
	IE   Word // 0xFFFF
}

func (mem *Memory) ResetMemory() {
	for i := range mem.Cart {
		mem.Cart[i] = 0
	}
	for i := range mem.VRAM {
		mem.VRAM[i] = 0
	}
	for i := range mem.ExRAM {
		mem.ExRAM[i] = 0
	}
	for i := range mem.WRAM {
		mem.WRAM[i] = 0
	}
	for i := range mem.OAM {
		mem.OAM[i] = 0
	}
	for i := range mem.IO {
		mem.IO[i] = 0
	}
	for i := range mem.HRAM {
		mem.HRAM[i] = 0
	}
	mem.IE = 0
}

func (mem *Memory) Read(address int) Word {
	switch {
	case address < 0x8000:
		return mem.Cart[address]
	case address < 0xA000:
		return mem.VRAM[address-0x8000]
	case address < 0xC000:
		return mem.ExRAM[address-0xA000]
	case address < 0xE000:
		return mem.WRAM[address-0xC000]
	case address < 0xFE00:
		return 0 // prohibited area
	case address < 0xFEA0:
		return mem.OAM[address-0xFE00]
	case address < 0xFF00:
		return 0 // unusable
	case address < 0xFF80:
		return mem.IO[address-0xFF00]
	case address < 0xFFFF:
		return mem.HRAM[address-0xFF80]
	case address == 0xFFFF:
		return mem.IE
	default:
		return 0
	}
}
