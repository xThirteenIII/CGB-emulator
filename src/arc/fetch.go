package arc

import "log"

// FetchByte reads the next byte pointed by PC, increases PC and consumes one cycle.
// It returns the byte read.
func (cpu *CPU) FetchByte(cycles *int) byte {

    // Exceeding max memory halts the cpu (Fatal log)
    if cpu.Registers.PC > MaxMem-1 {
        log.Fatalf("Program Counter exceeded max memory.")
    }

    // Fetch instruction at Program Counter address.
    byteRead := cpu.Memory.RAM[cpu.Registers.PC]

    // Increment Program Counter.
    cpu.Registers.PC++

    // Consume one clock cycle.
    *cycles--

    return byteRead
}

// FetchByte reads the next byte pointed by PC, increases PC + 2 and consumes two cycle.
// It returns the byte read.
func (cpu *CPU) FetchWord(cycles *int) uint16 {

    // Exceeding max memory halts the cpu (Fatal log)
    if cpu.Registers.PC > MaxMem-1 {
        log.Fatalf("Program Counter exceeded max memory.")
    }

    lsb := cpu.Memory.RAM[cpu.Registers.PC]
    cpu.Registers.PC++ 
    *cycles--

    msb := cpu.Memory.RAM[cpu.Registers.PC]
    cpu.Registers.PC++
    *cycles--

    word := uint16(msb) << 8 | uint16(lsb)

    return word
}
