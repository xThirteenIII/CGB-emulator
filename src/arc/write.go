package arc

import "log"

// WriteByteToMemory writes a byte into the absolute address location.
// It consumes one clock cycle.
func (cpu *CPU) WriteByteToMemory(cycles *int, address uint16, data byte) {
    if address > MaxMem - 1 {
        log.Fatalf("Address %d exceeded max memory.", address)
    }

    cpu.Memory.RAM[address] = data
    *cycles--
}
