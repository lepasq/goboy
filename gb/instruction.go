package gb

func (gb *Gameboy) popPC() byte {
	opcode := gb.Mem.Read(gb.CPU.PC.value)
	gb.CPU.PC.value++
	return opcode
}

func (gb *Gameboy) popPC16() uint16 {
	b1 := uint16(gb.popPC())
	b2 := uint16(gb.popPC())
	return b2<<8 | b1
}

var instructions = map[int]func(*Gameboy){
	0x06: func(gb *Gameboy) {
		// LD B, n
		gb.CPU.BC.SetHi(gb.popPC())
	},
	0x0E: func(gb *Gameboy) {
		// LD C, n
		gb.CPU.BC.SetLo(gb.popPC())
	},
	0x16: func(gb *Gameboy) {
		// LD D, n
		gb.CPU.DE.SetHi(gb.popPC())
	},
	0x1E: func(gb *Gameboy) {
		// LD E, n
		gb.CPU.DE.SetLo(gb.popPC())
	},
	0x26: func(gb *Gameboy) {
		// LD H, n
		gb.CPU.HL.SetHi(gb.popPC())
	},
	0x2E: func(gb *Gameboy) {
		// LD L, n
		gb.CPU.HL.SetLo(gb.popPC())
	},
	0x7F: func(gb *Gameboy) {
		// LD A, A
		gb.CPU.AF.SetHi(gb.CPU.AF.Hi())
	},
	0x78: func(gb *Gameboy) {
		// LD A,B
		gb.CPU.AF.SetHi(gb.CPU.BC.Hi())
	},
	0x79: func(gb *Gameboy) {
		// LD A,C
		gb.CPU.AF.SetHi(gb.CPU.BC.Lo())
	},
	0x7A: func(gb *Gameboy) {
		// LD A,D
		gb.CPU.AF.SetHi(gb.CPU.DE.Hi())
	},
	0x7B: func(gb *Gameboy) {
		// LD A,E
		gb.CPU.AF.SetHi(gb.CPU.DE.Lo())
	},
	0x7C: func(gb *Gameboy) {
		// LD A,H
		gb.CPU.AF.SetHi(gb.CPU.HL.Hi())
	},
	0x7D: func(gb *Gameboy) {
		// LD A,L
		gb.CPU.AF.SetHi(gb.CPU.HL.Lo())
	},
	0x0A: func(gb *Gameboy) {
		// LD A,(BC)
		val := gb.Mem.Read(gb.CPU.BC.HiLo())
		gb.CPU.AF.SetHi(val)
	},
	0x1A: func(gb *Gameboy) {
		// LD A,(DE)
		val := gb.Mem.Read(gb.CPU.DE.HiLo())
		gb.CPU.AF.SetHi(val)
	},
	0x7E: func(gb *Gameboy) {
		// LD A,(HL)
		val := gb.Mem.Read(gb.CPU.HL.HiLo())
		gb.CPU.AF.SetHi(val)
	},
	0xFA: func(gb *Gameboy) {
		// LD A,(nn)
		val := gb.Mem.Read(gb.popPC16())
		gb.CPU.AF.SetHi(val)
	},
	0x3E: func(gb *Gameboy) {
		// LD A,(nn)
		val := gb.popPC()
		gb.CPU.AF.SetHi(val)
	},
	0x47: func(gb *Gameboy) {
		// LD B,A
		gb.CPU.BC.SetHi(gb.CPU.AF.Hi())
	},
	0x40: func(gb *Gameboy) {
		// LD B,B
		gb.CPU.BC.SetHi(gb.CPU.BC.Hi())
	},
	0x41: func(gb *Gameboy) {
		// LD B,C
		gb.CPU.BC.SetHi(gb.CPU.BC.Lo())
	},
	0x42: func(gb *Gameboy) {
		// LD B,D
		gb.CPU.BC.SetHi(gb.CPU.DE.Hi())
	},
	0x43: func(gb *Gameboy) {
		// LD B,E
		gb.CPU.BC.SetHi(gb.CPU.DE.Lo())
	},
	0x44: func(gb *Gameboy) {
		// LD B,H
		gb.CPU.BC.SetHi(gb.CPU.HL.Hi())
	},
	0x45: func(gb *Gameboy) {
		// LD B,L
		gb.CPU.BC.SetHi(gb.CPU.HL.Lo())
	},
	0x46: func(gb *Gameboy) {
		// LD B,(HL)
		val := gb.Mem.Read(gb.CPU.HL.HiLo())
		gb.CPU.BC.SetHi(val)
	},
	0x4F: func(gb *Gameboy) {
		// LD C,A
		gb.CPU.BC.SetLo(gb.CPU.AF.Hi())
	},
	0x48: func(gb *Gameboy) {
		// LD C,B
		gb.CPU.BC.SetLo(gb.CPU.BC.Hi())
	},
	0x49: func(gb *Gameboy) {
		// LD C,C
		gb.CPU.BC.SetLo(gb.CPU.BC.Lo())
	},
	0x4A: func(gb *Gameboy) {
		// LD C,D
		gb.CPU.BC.SetLo(gb.CPU.DE.Hi())
	},
	0x4B: func(gb *Gameboy) {
		// LD C,E
		gb.CPU.BC.SetLo(gb.CPU.DE.Lo())
	},
	0x4C: func(gb *Gameboy) {
		// LD C,H
		gb.CPU.BC.SetLo(gb.CPU.HL.Hi())
	},
	0x4D: func(gb *Gameboy) {
		// LD C,L
		gb.CPU.BC.SetLo(gb.CPU.HL.Lo())
	},
	0x4E: func(gb *Gameboy) {
		// LD C,(HL)
		val := gb.Mem.Read(gb.CPU.HL.HiLo())
		gb.CPU.BC.SetLo(val)
	},
	0x57: func(gb *Gameboy) {
		// LD D,A
		gb.CPU.DE.SetHi(gb.CPU.AF.Hi())
	},
	0x50: func(gb *Gameboy) {
		// LD D,B
		gb.CPU.DE.SetHi(gb.CPU.BC.Hi())
	},
	0x51: func(gb *Gameboy) {
		// LD D,C
		gb.CPU.DE.SetHi(gb.CPU.BC.Lo())
	},
	0x52: func(gb *Gameboy) {
		// LD D,D
		gb.CPU.DE.SetHi(gb.CPU.DE.Hi())
	},
	0x53: func(gb *Gameboy) {
		// LD D,E
		gb.CPU.DE.SetHi(gb.CPU.DE.Lo())
	},
	0x54: func(gb *Gameboy) {
		// LD D,H
		gb.CPU.DE.SetHi(gb.CPU.HL.Hi())
	},
	0x55: func(gb *Gameboy) {
		// LD D,L
		gb.CPU.DE.SetHi(gb.CPU.HL.Lo())
	},
	0x56: func(gb *Gameboy) {
		// LD D,(HL)
		val := gb.Mem.Read(gb.CPU.HL.HiLo())
		gb.CPU.DE.SetHi(val)
	},
	0x5F: func(gb *Gameboy) {
		// LD E,A
		gb.CPU.DE.SetLo(gb.CPU.AF.Hi())
	},
	0x58: func(gb *Gameboy) {
		// LD E,B
		gb.CPU.DE.SetLo(gb.CPU.BC.Hi())
	},
	0x59: func(gb *Gameboy) {
		// LD E,C
		gb.CPU.DE.SetLo(gb.CPU.BC.Lo())
	},
	0x5A: func(gb *Gameboy) {
		// LD E,D
		gb.CPU.DE.SetLo(gb.CPU.DE.Hi())
	},
	0x5B: func(gb *Gameboy) {
		// LD E,E
		gb.CPU.DE.SetLo(gb.CPU.DE.Lo())
	},
	0x5C: func(gb *Gameboy) {
		// LD E,H
		gb.CPU.DE.SetLo(gb.CPU.HL.Hi())
	},
	0x5D: func(gb *Gameboy) {
		// LD E,L
		gb.CPU.DE.SetLo(gb.CPU.HL.Lo())
	},
	0x5E: func(gb *Gameboy) {
		// LD E,(HL)
		val := gb.Mem.Read(gb.CPU.HL.HiLo())
		gb.CPU.DE.SetLo(val)
	},
	0x67: func(gb *Gameboy) {
		// LD H,A
		gb.CPU.HL.SetHi(gb.CPU.AF.Hi())
	},
	0x60: func(gb *Gameboy) {
		// LD H,B
		gb.CPU.HL.SetHi(gb.CPU.BC.Hi())
	},
	0x61: func(gb *Gameboy) {
		// LD H,C
		gb.CPU.HL.SetHi(gb.CPU.BC.Lo())
	},
	0x62: func(gb *Gameboy) {
		// LD H,D
		gb.CPU.HL.SetHi(gb.CPU.DE.Hi())
	},
	0x63: func(gb *Gameboy) {
		// LD H,E
		gb.CPU.HL.SetHi(gb.CPU.DE.Lo())
	},
	0x64: func(gb *Gameboy) {
		// LD H,H
		gb.CPU.HL.SetHi(gb.CPU.HL.Hi())
	},
	0x65: func(gb *Gameboy) {
		// LD H,L
		gb.CPU.HL.SetHi(gb.CPU.HL.Lo())
	},
	0x66: func(gb *Gameboy) {
		// LD H,(HL)
		val := gb.Mem.Read(gb.CPU.HL.HiLo())
		gb.CPU.HL.SetHi(val)
	},
	0x6F: func(gb *Gameboy) {
		// LD L,A
		gb.CPU.HL.SetLo(gb.CPU.AF.Hi())
	},
	0x68: func(gb *Gameboy) {
		// LD L,B
		gb.CPU.HL.SetLo(gb.CPU.BC.Hi())
	},
	0x69: func(gb *Gameboy) {
		// LD L,C
		gb.CPU.HL.SetLo(gb.CPU.BC.Lo())
	},
	0x6A: func(gb *Gameboy) {
		// LD L,D
		gb.CPU.HL.SetLo(gb.CPU.DE.Hi())
	},
	0x6B: func(gb *Gameboy) {
		// LD L,E
		gb.CPU.HL.SetLo(gb.CPU.DE.Lo())
	},
	0x6C: func(gb *Gameboy) {
		// LD L,H
		gb.CPU.HL.SetLo(gb.CPU.HL.Hi())
	},
	0x6D: func(gb *Gameboy) {
		// LD L,L
		gb.CPU.HL.SetLo(gb.CPU.HL.Lo())
	},
	0x6E: func(gb *Gameboy) {
		// LD L,(HL)
		val := gb.Mem.Read(gb.CPU.HL.HiLo())
		gb.CPU.HL.SetLo(val)
	},
	0x77: func(gb *Gameboy) {
		// LD (HL),A
		val := gb.CPU.AF.Hi()
		gb.Mem.Write(gb.CPU.HL.HiLo(), val)
	},
	0x70: func(gb *Gameboy) {
		// LD (HL),B
		val := gb.CPU.BC.Hi()
		gb.Mem.Write(gb.CPU.HL.HiLo(), val)
	},
	0x71: func(gb *Gameboy) {
		// LD (HL),C
		val := gb.CPU.BC.Lo()
		gb.Mem.Write(gb.CPU.HL.HiLo(), val)
	},
	0x72: func(gb *Gameboy) {
		// LD (HL),D
		val := gb.CPU.DE.Hi()
		gb.Mem.Write(gb.CPU.HL.HiLo(), val)
	},
	0x73: func(gb *Gameboy) {
		// LD (HL),E
		val := gb.CPU.DE.Lo()
		gb.Mem.Write(gb.CPU.HL.HiLo(), val)
	},
	0x74: func(gb *Gameboy) {
		// LD (HL),H
		val := gb.CPU.HL.Hi()
		gb.Mem.Write(gb.CPU.HL.HiLo(), val)
	},
	0x75: func(gb *Gameboy) {
		// LD (HL),L
		val := gb.CPU.HL.Lo()
		gb.Mem.Write(gb.CPU.HL.HiLo(), val)
	},
	0x36: func(gb *Gameboy) {
		// LD (HL),n 36
		val := gb.popPC()
		gb.Mem.Write(gb.CPU.HL.HiLo(), val)
	},
	0x02: func(gb *Gameboy) {
		// LD (BC),A
		val := gb.CPU.AF.Hi()
		gb.Mem.Write(gb.CPU.BC.HiLo(), val)
	},
	0x12: func(gb *Gameboy) {
		// LD (DE),A
		val := gb.CPU.AF.Hi()
		gb.Mem.Write(gb.CPU.DE.HiLo(), val)
	},
	0xEA: func(gb *Gameboy) {
		// LD (nn),A
		val := gb.CPU.AF.Hi()
		gb.Mem.Write(gb.popPC16(), val)
	},
	0xF2: func(gb *Gameboy) {
		// LD A,(C)
		val := 0xFF00 + uint16(gb.CPU.BC.Lo())
		gb.CPU.AF.SetHi(gb.Mem.Read(val))
	},
	0xE2: func(gb *Gameboy) {
		// LD (C),A
		val := gb.CPU.AF.Hi()
		mem := 0xFF00 + uint16(gb.CPU.BC.Lo())
		gb.Mem.Write(mem, val)
	},
	0x3A: func(gb *Gameboy) {
		// LDD A,(HL)
		val := gb.Mem.Read(gb.CPU.HL.HiLo())
		gb.CPU.AF.SetHi(val)
		gb.CPU.HL.Set(gb.CPU.HL.HiLo() - 1)
	},
	0x32: func(gb *Gameboy) {
		// LDD (HL),A
		val := gb.CPU.HL.HiLo()
		gb.Mem.Write(val, gb.CPU.AF.Hi())
		gb.CPU.HL.Set(gb.CPU.HL.HiLo() - 1)
	},
	0x2A: func(gb *Gameboy) {
		// LDI A,(HL)
		val := gb.Mem.Read(gb.CPU.HL.HiLo())
		gb.CPU.AF.SetHi(val)
		gb.CPU.HL.Set(gb.CPU.HL.HiLo() + 1)
	},
	0x22: func(gb *Gameboy) {
		// LDI (HL),A
		val := gb.CPU.HL.HiLo()
		gb.Mem.Write(val, gb.CPU.AF.Hi())
		gb.CPU.HL.Set(gb.CPU.HL.HiLo() + 1)
	},
	0xE0: func(gb *Gameboy) {
		// LD (0xFF00+n),A
		val := 0xFF00 + uint16(gb.popPC())
		gb.Mem.Write(val, gb.CPU.AF.Hi())
	},
	0xF0: func(gb *Gameboy) {
		// LD A,(0xFF00+n)
		val := gb.Mem.Read(0xFF00 + uint16(gb.popPC())) // TODO: fix?
		gb.CPU.AF.SetHi(val)
	},
	// ========== 16-Bit Loads ===========
	0x01: func(gb *Gameboy) {
		// LD BC,nn
		val := gb.popPC16()
		gb.CPU.BC.Set(val)
	},
	0x11: func(gb *Gameboy) {
		// LD DE,nn
		val := gb.popPC16()
		gb.CPU.DE.Set(val)
	},
	0x21: func(gb *Gameboy) {
		// LD HL,nn
		val := gb.popPC16()
		gb.CPU.HL.Set(val)
	},
	0x31: func(gb *Gameboy) {
		// LD SP,nn
		val := gb.popPC16()
		gb.CPU.SP.Set(val)
	},
	0xF9: func(gb *Gameboy) {
		// LD SP,HL
		val := gb.CPU.HL
		gb.CPU.SP = val
	},
}
