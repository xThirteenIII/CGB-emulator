package arc

import "log"

// ReadByteFromMemory returns the byte read at the address location in memory.
// It consumes one clock cycle, without increasing the PC.
func (cpu *CPU) ReadByteFromMemory(cycles *int, address uint16) byte{

    if address > MaxMem - 1 {
        log.Fatalf("PC exceeded max memory.")
    }

    byteRead := cpu.Memory.RAM[cpu.Registers.PC]
    *cycles--


    return byteRead
}
