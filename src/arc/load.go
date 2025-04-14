package arc

// PopFromSP returns the contents from the memory stack and increments SP by 1.
// It consumes 1 machine cycle.
// WARNING: SP might go into safe area (>0xFFFE) && (< 0xC000), might need a check later.
func (cpu *CPU) PopFromSP(cycles *int) byte {

    data := cpu.Memory.RAM[cpu.Registers.SP]
    cpu.Registers.SP++
    *cycles--

    return data
}

