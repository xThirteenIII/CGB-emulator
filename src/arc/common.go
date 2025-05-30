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

// AddInt8ToUint16WithoutCarry adds the signed 8-bit byte to the unsigned uint16 value.
// It returns the result of the operation and a byte with Carry and HalfCarry set according 
// to operation result.
func AddInt8ToUint16WithoutCarry(value uint16, adder int8) (uint16, byte) {

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

// AddByteToByteWithoutCarry adds b1 to b.
// It returns the result of the operation and a byte with Carry and HalfCarry bits set according 
// to operation result.
func AddByteToByteWithCarry(b , b1, flags byte) (byte, byte) {

    carryPerBit := byte(0)
    var carryFlag byte
    if (flags & (1 << 4)) == 0 {
        carryFlag = 0x00
    }else {
        carryFlag = 0x01
    }

    // Isolate four lower bits of the 16bit address and the signed 8bit.
    if ((b&0x0F) + (b1 & 0x0F) + carryFlag) > 0x0F {

        // Set halfCarryBit, which is bit 5
        carryPerBit |= 1 << 5
    }

    if uint16(b) + uint16(b1) + uint16(carryFlag)> 0xFF {

        // Set CarryBit, which is bit 4
        carryPerBit |= 1 << 4 // that is: | 0b00010000
    }

    return b + b1 + carryFlag, carryPerBit
}

// AddByteToByteWithoutCarry adds b1 to b.
// It returns the result of the operation and a byte with Carry and HalfCarry bits set according 
// to operation result.
func AddByteToByteWithoutCarry(b , b1 byte) (byte, byte) {

    carryPerBit := byte(0)

    // Isolate four lower bits of the 16bit address and the signed 8bit.
    if ((b&0x0F) + (b1 & 0x0F)) > 0x0F {

        // Set halfCarryBit, which is bit 5
        carryPerBit |= 1 << 5
    }

    if uint16(b) + uint16(b1) > 0xFF {

        // Set CarryBit, which is bit 4
        carryPerBit |= 1 << 4 // that is: | 0b00010000
    }

    return b + b1, carryPerBit
}

// AddByteToByteWithoutCarry adds b1 to b.
// It returns the result of the operation and a byte with Carry and HalfCarry bits set according 
// to operation result.
func AddWordToWordWithoutCarry(w , w1 uint16) (uint16, byte) {

    carryPerBit := byte(0)

    // Isolate 12 lower bits of the 16bit address and the signed 8bit.
    if ((w&0x0FFF) + (w1 & 0x0FFF)) > 0x0FFF {

        // Set halfCarryBit, which is bit 5
        carryPerBit |= 1 << 5
    }

    if uint32(w) + uint32(w1) > 0xFFFF {

        // Set CarryBit, which is bit 4
        carryPerBit |= 1 << 4 // that is: | 0b00010000
    }

    return w + w1, carryPerBit
}

// SubByteFromByteWithCarry subtracts b1 from b.
// It returns the result of the operation and a byte with Carry and HalfCarry bits set according 
// to operation result.
func SubByteFromByteWithCarry(b , b1, flags byte) (byte, byte) {

    carryPerBit := byte(0)
    var carryFlag byte
    if (flags & (1 << 4)) == 0 {
        carryFlag = 0x00
    }else {
        carryFlag = 0x01
    }

    if  (b&0x0F) < ((b1&0x0F) + carryFlag) {

        // Set halfCarryBit, which is bit 5
        carryPerBit |= 1 << 5
    }

    if  b < (b1 + carryFlag) {

        // Set CarryBit, which is bit 4
        carryPerBit |= 1 << 4 // that is: | 0b00010000
    }

    return b - b1 - carryFlag, carryPerBit
}

// SubByteFromByteWithoutCarry subtracts b1 from b.
// It returns the result of the operation and a byte with Carry and HalfCarry bits set according 
// to operation result.
func SubByteFromByteWithoutCarry(b , b1 byte) (byte, byte) {

    carryPerBit := byte(0)

    // If lower nibble of b1 is higher than lower nibble of b, a borrow is happening.
    if (b&0x0F) < (b1&0x0F) {

        // Set halfCarryBit, which is bit 5
        carryPerBit |= 1 << 5
    }

    if b < b1 {

        // Set CarryBit, which is bit 4
        carryPerBit |= 1 << 4 // that is: | 0b00010000
    }

    return b - b1, carryPerBit
}

// IncrementByteBy1 adds 1 to the byte.
// It returns the result of the operation and a bool set to true if half carry happens.
func IncrementByteBy1(value byte) (byte, bool) {

    return value + 1, (value & 0x0F) + 1 > 0x0F
}

// DecrementByteBy1 subtracts 1 to the byte.
// Half Carry happens when the lower 4 bits need to borrow 1 from the upper 4 bits.
// E.g. 0x10 - 1 = 0x0F -> half carry set
func DecrementByteBy1(value byte) (byte, bool) {

    return value - 1, (value & 0x0F) == 0 // Half Carry happens only if value has the last 4 bits set to 0.
}

func (cpu *CPU) SetZflag() {
    
    cpu.Registers.F |= 1 << 7
}

func (cpu *CPU) SetNflag() {
    
    cpu.Registers.F |= 1 << 6
}

func (cpu *CPU) SetHflag() {
    
    cpu.Registers.F |= 1 << 5
}

func (cpu *CPU) SetCflag() {
    
    cpu.Registers.F |= 1 << 4
}

func (cpu *CPU) ClearZflag() {
    
    cpu.Registers.F &^= 1 << 7
}

func (cpu *CPU) ClearNflag() {
    
    cpu.Registers.F &^= 1 << 6
}

func (cpu *CPU) ClearHflag() {
    
    cpu.Registers.F &^= 1 << 5
}

func (cpu *CPU) ClearCflag() {
    
    cpu.Registers.F &^= 1 << 4
}

