package gb

import "fmt"

type Byte = uint8
type Word = uint16

type CPU struct {
	AF register /** 7: zero, 6: subtraction, 5 half carry, 4: carry */
	BC register
	DE register
	HL register

	PC register
	SP register
}

type register struct {
	value Word
}

func (cpu *CPU) String() string {
	return fmt.Sprintf(`
    A: %b      F: %b
    B: %b      C: %b
    D: %b      E: %b
    H: %b      L: %b
    
    PC: %b
    SP: %b`, cpu.AF.Hi(), cpu.AF.Lo(), cpu.BC.Hi(), cpu.BC.Lo(), cpu.DE.Hi(), cpu.DE.Lo(), cpu.HL.Hi(), cpu.HL.Lo(), cpu.PC, cpu.SP)
}

func (cpu *CPU) ResetRegisters() {
	cpu.AF.Set(0x1180)
	cpu.BC.Set(0)
	cpu.DE.Set(0xFF56)
	cpu.HL.Set(0x000D)

	cpu.PC.Set(0x100)
	cpu.SP.Set(0xFFFE)
}

func (cpu *CPU) Execute(cycles Word) {
	for cycles > 0 {
		// actually run instruction?
		cpu.PC.value += 1
		cycles--
	}
}

func (reg *register) setFlagHigh(position int) {
	var mask Byte = 1 << position
	reg.SetLo((reg.Lo() & ^mask) | (1 << position))
}

func (reg *register) setFlagLow(position int) {
	var mask Byte = 1 << position
	reg.SetLo((reg.Lo() & ^mask) | (0 << position))
}

func (reg *register) setFlag(position int, value bool) {
	if value {
		reg.setFlagHigh(position)
	} else {
		reg.setFlagLow(position)
	}
}

func (cpu *CPU) SetZero(value bool) {
	cpu.AF.setFlag(7, value)
}

func (cpu *CPU) SetSub(value bool) { // necessary if last operation was a subtraction
	cpu.AF.setFlag(6, value)
}

func (cpu *CPU) SetHalf(value bool) { // Set if, in the result of the last operation, the lower half of the byte overflowed past 15
	cpu.AF.setFlag(5, value)
}

func (cpu *CPU) SetCarry(value bool) { // Set if the last operation produced a result over 255 (for additions) or under 0 (for subtractions)
	cpu.AF.setFlag(4, value)
}

func (cpu *CPU) Z() bool {
	return cpu.AF.Lo()>>7&1 == 1
}

func (cpu *CPU) N() bool {
	return cpu.AF.Lo()>>6&1 == 1
}

func (cpu *CPU) H() bool {
	return cpu.AF.Lo()>>5&1 == 1
}

func (cpu *CPU) C() bool {
	return cpu.AF.Lo()>>4&1 == 1
}

func (reg *register) Hi() Byte {
	return byte(reg.value >> 8)
}

func (reg *register) Lo() Byte {
	return byte(reg.value & 0xFF)
}

func (reg *register) HiLo() Word {
	return reg.value
}

func (reg *register) Set(newValue Word) {
	reg.value = newValue
}

func (reg *register) SetHi(newValue Byte) {
	reg.value = uint16(newValue)<<8 | (uint16(reg.value) & 0xFF)
}

func (reg *register) SetLo(newValue Byte) {
	reg.value = uint16(newValue) | (uint16(reg.value) & 0xFF00)
}

func ExampleProgram() {
	var cpu CPU
	cpu.ResetRegisters()

	cpu.SetZero(true)
	cpu.SetCarry(true)
	fmt.Println(&cpu)

	fmt.Println(cpu.Z(), cpu.N(), cpu.H(), cpu.C())
}
