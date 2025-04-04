package arc

import (
	"cgbemu/src/instructions"
	"testing"
)

func TestLDBImmediate(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDB_IM
    cpu.Memory.RAM[0x0101] = 0xF2

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.B != 0xF2 {
        t.Error("B register should be 0xF2, instead got: ", cpu.Registers.B)
    }
}

func TestLDB_HL(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x80
    cpu.Registers.L = 0x8F
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDB_HL
    cpu.Memory.RAM[0x808F] = 0x20

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.B != 0x20 {
        t.Error("B register should be 0x20, instead got: ", cpu.Registers.B)
    }
}
