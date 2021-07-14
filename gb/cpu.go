package gb

import "fmt"

type Byte = uint8
type Word = uint16

type CPU struct {
	af Word
	bc Word
	de Word
	hl Word

	pc Word
	sp Word

	z Byte // zero flag
	n Byte // subtraction flag
	h Byte // half carry flag
	c Byte // carry flag
}

func (cpu CPU) String() string {
	return fmt.Sprintf(`
    af: %d
    bc: %d
    de: %d
    hl: %d
    
    pc: %d
    sp: %d
    
    z: %d
    n: %d
    h: %d
    c: %d`,
		cpu.af, cpu.bc, cpu.de, cpu.hl, cpu.pc, cpu.sp, cpu.z, cpu.n, cpu.h, cpu.c)
}

func (cpu CPU) ResetRegisters() {
	cpu.af = 0
	cpu.bc = 0
	cpu.de = 0
	cpu.hl = 0
	cpu.pc = 0
	cpu.sp = 0
	cpu.z = 0
	cpu.n = 0
	cpu.h = 0
	cpu.c = 0
}
