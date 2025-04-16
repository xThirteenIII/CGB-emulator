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

func TestINC_C(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.C = 0x34
    cpu.Memory.RAM[0x0100] = instructions.INC_C

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.C != 0x35 {
        t.Error("C register should be 0x35, instead got: ", cpu.Registers.C)
    }
}

func TestINC_CSetsZandHFlags(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.C = 0xFF
    cpu.Memory.RAM[0x0100] = instructions.INC_C

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.C != 0x00 {
        t.Error("C register should be 0, instead got: ", cpu.Registers.C)
    }

    if cpu.Registers.F != 0b10100000 {
        t.Error("Z and H should be set. Instead got: ", cpu.Registers.F)
    }
}

func TestDEC_C(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.C = 0x34
    cpu.Memory.RAM[0x0100] = instructions.DEC_C

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.C != 0x33 {
        t.Error("C register should be 0x33, instead got: ", cpu.Registers.C)
    }
}

func TestDEC_CSetsZandHFlags(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.C = 0xF0
    cpu.Memory.RAM[0x0100] = instructions.DEC_C

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.C != 0xEF {
        t.Error("C register should be 0xEF, instead got: ", cpu.Registers.C)
    }

    if cpu.Registers.F != 0b01100000 {
        t.Error("Z, N and H should be set. Instead got: ", cpu.Registers.F)
    }
}

func TestDEC_E(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.E = 0x34
    cpu.Memory.RAM[0x0100] = instructions.DEC_E

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.E != 0x33 {
        t.Error("E register should be 0x33, instead got: ", cpu.Registers.E)
    }
}

func TestDEC_ESetsZandHFlags(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.E = 0xF0
    cpu.Memory.RAM[0x0100] = instructions.DEC_E

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.E != 0xEF {
        t.Error("E register should be 0xEF, instead got: ", cpu.Registers.E)
    }

    if cpu.Registers.F != 0b01100000 {
        t.Error("Z, N and H should be set. Instead got: ", cpu.Registers.F)
    }
}

func TestINC_E(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.E = 0x34
    cpu.Memory.RAM[0x0100] = instructions.INC_E

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.E != 0x35 {
        t.Error("E register should be 0x35, instead got: ", cpu.Registers.E)
    }
}

func TestINC_ESetsZandHFlags(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.E = 0xFF
    cpu.Memory.RAM[0x0100] = instructions.INC_E

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.E != 0x00 {
        t.Error("E register should be 0, instead got: ", cpu.Registers.E)
    }

    if cpu.Registers.F != 0b10100000 {
        t.Error("Z and H should be set. Instead got: ", cpu.Registers.F)
    }
}

func TestINC_H(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.H = 0x34
    cpu.Memory.RAM[0x0100] = instructions.INC_H

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.H != 0x35 {
        t.Error("H register should be 0x35, instead got: ", cpu.Registers.H)
    }
}

func TestINC_HSetsZandHFlags(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.H = 0xFF
    cpu.Memory.RAM[0x0100] = instructions.INC_H

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.H != 0x00 {
        t.Error("H register should be 0, instead got: ", cpu.Registers.H)
    }

    if cpu.Registers.F != 0b10100000 {
        t.Error("Z and H should be set. Instead got: ", cpu.Registers.F)
    }
}

func TestDEC_H(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.H = 0x34
    cpu.Memory.RAM[0x0100] = instructions.DEC_H

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.H != 0x33 {
        t.Error("H register should be 0x33, instead got: ", cpu.Registers.H)
    }
}

func TestDEC_HSetsZandHFlags(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.H = 0xF0
    cpu.Memory.RAM[0x0100] = instructions.DEC_H

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.H != 0xEF {
        t.Error("H register should be 0xEF, instead got: ", cpu.Registers.H)
    }

    if cpu.Registers.F != 0b01100000 {
        t.Error("Z, N and H should be set. Instead got: ", cpu.Registers.F)
    }
}
