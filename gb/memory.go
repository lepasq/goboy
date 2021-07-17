package gb

type Memory struct { // make this a map, with the starting address as a value
	Cart  [0x8000]byte
	VRAM  [0x2000]byte
	ExRAM [0x2000]byte
	WRAM  [0x1000]byte
	// -- prohibited --
	OAM [0xA0]byte
	// -- unusable --
	IO   [0x80]byte
	HRAM [0x7E]byte
	IE   byte // 0xFFFF
}

type Section struct {
	start int
	end   int
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

func (mem *Memory) Read(address int) byte {
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

func (mem *Memory) Write(address int, value byte) {
	switch {
	case address < 0x8000:
		mem.Cart[address] = value
		break
	case address < 0xA000:
		mem.VRAM[address-0x8000] = value
		break
	case address < 0xC000:
		mem.ExRAM[address-0xA000] = value
		break
	case address < 0xE000:
		mem.WRAM[address-0xC000] = value
		break
	case address < 0xFE00:
		break
	case address < 0xFEA0:
		mem.OAM[address-0xFE00] = value
		break
	case address < 0xFF00:
		break
	case address < 0xFF80:
		mem.IO[address-0xFF00] = value
		break
	case address < 0xFFFF:
		mem.HRAM[address-0xFF80] = value
		break
	case address == 0xFFFF:
		mem.IE = value
		break
	}
}
