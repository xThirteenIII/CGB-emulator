package arc

import (
	"cgbemu/src/instructions"
	"testing"
)

func TestINC_B(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.B = 0x34
    cpu.Memory.RAM[0x0100] = instructions.INC_B

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.B != 0x35 {
        t.Error("B register should be 0x35, instead got: ", cpu.Registers.B)
    }
}

func TestINC_BSetsZandHFlags(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.B = 0xFF
    cpu.Memory.RAM[0x0100] = instructions.INC_B

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.B != 0x00 {
        t.Error("B register should be 0, instead got: ", cpu.Registers.B)
    }

    if cpu.Registers.F != 0b10100000 {
        t.Error("Z and H should be set. Instead got: ", cpu.Registers.F)
    }
}

func TestDEC_B(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.B = 0x34
    cpu.Memory.RAM[0x0100] = instructions.DEC_B

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.B != 0x33 {
        t.Error("B register should be 0x33, instead got: ", cpu.Registers.B)
    }
}

func TestDEC_BSetsZandHFlags(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.B = 0xF0
    cpu.Memory.RAM[0x0100] = instructions.DEC_B

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.B != 0xEF {
        t.Error("B register should be 0xEF, instead got: ", cpu.Registers.B)
    }

    if cpu.Registers.F != 0b01100000 {
        t.Error("Z, N and H should be set. Instead got: ", cpu.Registers.F)
    }
}
