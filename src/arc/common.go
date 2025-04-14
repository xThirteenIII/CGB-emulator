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

// AddInt8ToUint16 adds the signed 8-bit byte to the unsigned uint16. 
// It returns the result of the operation and a byte with Carry and HalfCarry set according 
// to operation result.
func AddInt8ToUint16WithCarry(value uint16, adder int8) (uint16, byte) {

    carryPerBit := byte(0)

    // Isolate four lower bits of the 16bit address and the signed 8bit.
    if ((value&0x0F) + (uint16(adder) & 0x0F)) > 0x0F {

        // Set halfCarryBit, which is bit 5
        carryPerBit |= 1 << 5
    }

    if ((value & 0xFF) + (uint16(adder) & 0xFF)) > 0xFF {

        // Set CarryBit, which is bit 4
        carryPerBit |= 1 << 4 // that is: | 0b00010000
    }

    return uint16(int32(value) + int32(adder)), carryPerBit
}

