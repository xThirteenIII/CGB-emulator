package arc

import "log"

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
