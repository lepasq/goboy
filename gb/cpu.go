package gb

import "fmt"

type Byte = uint8
type Word = uint16

type CPU struct {
	A Byte
	F Byte /** 7: zero, 6: subtraction, 5 half carry, 4: carry */
	B Byte
	C Byte
	D Byte
	E Byte
	H Byte
	L Byte

	PC Word
	SP Word
}

func (cpu *CPU) String() string {
	return fmt.Sprintf(`
    A: %b      F: %b
    B: %b      C: %b
    D: %b      E: %b
    H: %b      L: %b
    
    PC: %b
    SP: %b`, cpu.A, cpu.F, cpu.B, cpu.C, cpu.D, cpu.E, cpu.H, cpu.L, cpu.PC, cpu.SP)
}

func (cpu *CPU) ResetRegisters() {
	cpu.A = 0x11
	cpu.F = 0x80
	cpu.B = 0
	cpu.C = 0
	cpu.D = 0xFF
	cpu.E = 0x56
	cpu.H = 0
	cpu.L = 0x0D

	cpu.PC = 0x100
	cpu.SP = 0xFFFE
}

func (cpu *CPU) Execute(cycles Word) {
	for cycles > 0 {
		// actually run instruction?
		cpu.PC += 1
		cycles--
	}
}

func (cpu *CPU) setHigh(position int) {
	position -= 1
	var mask uint8 = 1 << position
	cpu.F = ((cpu.F & ^mask) | (1 << position))
}

func (cpu *CPU) setLow(position int) {
	position -= 1
	var mask uint8 = 1 << position
	cpu.F = ((cpu.F & ^mask) | (0 << position))
}

func (cpu *CPU) setFlag(position int, value bool) {
	if value {
		cpu.setHigh(position)
	} else {
		cpu.setLow(position)
	}
}

func (cpu *CPU) SetZero(value bool) {
	cpu.setFlag(7, value)
}

func (cpu *CPU) SetSub(value bool) { // necessary if last operation was a subtraction
	cpu.setFlag(6, value)
}

func (cpu *CPU) SetHalf(value bool) { // Set if, in the result of the last operation, the lower half of the byte overflowed past 15
	cpu.setFlag(5, value)
}

func (cpu *CPU) SetCarry(value bool) { // Set if the last operation produced a result over 255 (for additions) or under 0 (for subtractions)
	cpu.setFlag(4, value)
}

func ExampleProgram() {
	var cpu CPU
	cpu.ResetRegisters()

	cpu.SetZero(true)
	cpu.SetCarry(true)
	fmt.Println(&cpu)
}
