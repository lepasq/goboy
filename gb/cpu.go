package gb

import "fmt"

type Byte = uint8
type Word = uint16

type CPU struct {
	a Byte
	f Byte
	b Byte
	c Byte
	d Byte
	e Byte
	h Byte
	l Byte

	pc Word
	sp Word

	z     Byte // zero flag
	n     Byte // subtraction flag
	half  Byte // half carry flag
	carry Byte // carry flag
}

func (cpu *CPU) String() string {
	return fmt.Sprintf(`
    a: %d
    f: %d
    b: %d
    c: %d
    d: %d
    e: %d
    h: %d
    l: %d
    
    pc: %d
    sp: %d
    
    z: %d
    n: %d
    h: %d
    c: %d`,
		cpu.a, cpu.f, cpu.b, cpu.c, cpu.d, cpu.e, cpu.h, cpu.l, cpu.pc, cpu.sp, cpu.z, cpu.n, cpu.h, cpu.c)
}

func (cpu *CPU) ResetRegisters() {
	cpu.a = 0x11
	cpu.f = 0x80
	cpu.b = 0
	cpu.c = 0
	cpu.d = 0xFF
	cpu.e = 0x56
	cpu.h = 0
	cpu.l = 0x0D

	cpu.pc = 0x100
	cpu.sp = 0xFFFE

	cpu.z = 0
	cpu.n = 0
	cpu.h = 0
	cpu.c = 0
}

func (cpu *CPU) Execute(cycles Word) {
	for cycles > 0 {
		// actually run instruction?
		cpu.pc += 1
		cycles--
	}
}

func (cpu *CPU) setZero() {
}

func (cpu *CPU) setOp() { // necessary if last operation was a subtraction
}

func (cpu *CPU) setHalf() { // Set if, in the result of the last operation, the lower half of the byte overflowed past 15
}

func (cpu *CPU) SetCarry() { // Set if the last operation produced a result over 255 (for additions) or under 0 (for subtractions)
}

func (cpu *CPU) setFlag(flag Byte, value bool) {
	// if value {
	// cpu.flag = 1
	// } else {
	// cpu.flag = 0
	// }
}

func ExampleProgram() {
}
