package arc

// HL returns the HL content in 16 bit.
func (cpu *CPU) HL() uint16 {
    return uint16(cpu.Registers.H) << 8 | uint16(cpu.Registers.L)
}

// HL returns the BC content in 16 bit.
func (cpu *CPU) BC() uint16 {
    return uint16(cpu.Registers.B) << 8 | uint16(cpu.Registers.C)
}

// DE returns the DE content in 16 bit.
func (cpu *CPU) DE() uint16 {
    return uint16(cpu.Registers.D) << 8 | uint16(cpu.Registers.E)
}

// PopFromSP returns the contents from the memory stack and increments SP by 1.
// It consumes 1 machine cycle.
// WARNING: SP might go into safe area (>0xFFFE) && (< 0xC000), might need a check later.
func (cpu *CPU) PopFromSP(cycles *int) byte {

    data := cpu.Memory.RAM[cpu.Registers.SP]
    cpu.Registers.SP++
    *cycles--

    return data
}

func GetUint16AddressFromLSBAndMSB(lsb, msb byte) uint16 {

    return uint16(msb) << 8 | uint16(lsb)
}

// Increment16Address increments absolute address, given lsb and msb.
// TODO: Does HL=0xFFFF + 1 = 0000? Or need to handle exception?
func Increment16Address(lsb , msb *byte) {

    absoluteAddress := (uint16(*msb) << 8) | uint16(*lsb)
    absoluteAddress++

    *msb = byte(absoluteAddress >> 8)
    *lsb = byte(absoluteAddress & 0xFF)
}

// Increment16Address increments absolute address, given lsb and msb.
func Decrement16Address(lsb , msb *byte) {

    absoluteAddress := (uint16(*msb) << 8) | uint16(*lsb)
    absoluteAddress--

    *msb = byte(absoluteAddress >> 8)
    *lsb = byte(absoluteAddress & 0xFF)
}
