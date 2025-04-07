package arc

import "log"

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

func GetUint16AddressFromLSBAndMSB(lsb, msb byte) uint16 {

    return uint16(msb) << 8 | uint16(lsb)
}

// Increment16Address increments absolute address, given lsb and msb.
func Increment16Address(lsb , msb *byte) {

    absoluteAddress := (uint16(*msb) << 8) | uint16(*lsb)
    if absoluteAddress == 0xFFFF {
        log.Fatalf("16-bit register overflowed")
    }
    absoluteAddress++

    *msb = byte(absoluteAddress >> 8)
    *lsb = byte(absoluteAddress & 0xFF)
}

// Increment16Address increments absolute address, given lsb and msb.
func Decrement16Address(lsb , msb *byte) {

    absoluteAddress := (uint16(*msb) << 8) | uint16(*lsb)
    if absoluteAddress == 0x0000 {
        log.Fatalf("16-bit register overflowed")
    }
    absoluteAddress--

    *msb = byte(absoluteAddress >> 8)
    *lsb = byte(absoluteAddress & 0xFF)
}
