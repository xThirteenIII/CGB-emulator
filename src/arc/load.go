package arc

// 8-bit load instructions

// LDr8_n8 loads to the 8bit register r, the immediate data n.
func (cpu *CPU) LDr8_n8(register *byte, n byte) {
   *register = n 
}

func (cpu *CPU) LDr8_HL(register *byte) {
   *register = cpu.Memory.RAM[cpu.HL()]
   cpu.Registers.PC++
}

func (cpu *CPU) LDr8_r8(loadFrom, loadInto *byte) {
    *loadInto = *loadFrom
}

// HL returns the HL address in 16 bit.
func (cpu *CPU) HL() uint16 {
    return uint16(cpu.Registers.H) << 8 | uint16(cpu.Registers.L)
}
