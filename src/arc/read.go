package arc

import "log"

// ReadByteFromMemory returns the byte read at the address location in memory.
// It consumes one clock cycle, without increasing the PC.
func (cpu *CPU) ReadByteFromMemory(cycles *int, address uint16) byte{

    if address > MaxMem - 1 {
        log.Fatalf("PC exceeded max memory.")
    }

    byteRead := cpu.Memory.RAM[address]
    *cycles--


    return byteRead
}

// ReadByteFromMemory returns the word read at the address location in memory.
// It consumes one clock cycle, without increasing the PC.
func (cpu *CPU) ReadWordFromMemory(cycles *int, address uint16) uint16{

    if address > MaxMem - 1 {
        log.Fatalf("Address %d exceeded max memory.", address)
    }

    lsb := cpu.Memory.RAM[address]
    *cycles--

    msb := cpu.Memory.RAM[address+1]
    *cycles--

    return uint16(msb) << 8 | uint16(lsb)
}
