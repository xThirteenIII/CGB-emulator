package arc

// HL returns the HL address in 16 bit.
func (cpu *CPU) HL() uint16 {
    return uint16(cpu.Registers.H) << 8 | uint16(cpu.Registers.L)
}

// HL returns the HL address in 16 bit.
func (cpu *CPU) BC() uint16 {
    return uint16(cpu.Registers.B) << 8 | uint16(cpu.Registers.C)
}

func GetUint16AddressFromLSBAndMSB(lsb, msb byte) uint16 {

    return uint16(msb) << 8 | uint16(lsb)
}
