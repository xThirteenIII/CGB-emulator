package arc

import (
	"cgbemu/src/instructions"
	"testing"
)

func TestADDHL_BC(t *testing.T) {

    // Given
    cpu := InitSM83() 

    // When
    cpu.Registers.H = 0x13
    cpu.Registers.L = 0x20
    cpu.Registers.B = 0x25
    cpu.Registers.C = 0x2F
    cpu.Memory.RAM[0x0100] = instructions.ADDHL_BC

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.F & (1 << 6) != 0 {
        t.Error("N flag should be 0.")
    }

    if cpu.Registers.F & (1 << 5) != 0 {
        t.Error("H flag should be 0.")
    }

    if cpu.Registers.F & (1 << 4) != 0 {
        t.Error("C flag should be 0.")
    }

    if cpu.HL() != 0x384F {
        t.Error("HL should be 0x384F. Instead got: ", cpu.HL())
    }
}

func TestADDHL_BCSetsHandCFlags(t *testing.T) {

    // Given
    cpu := InitSM83() 

    // When
    cpu.Registers.H = 0x23
    cpu.Registers.L = 0x20
    cpu.Registers.B = 0xEF
    cpu.Registers.C = 0x2F
    cpu.Memory.RAM[0x0100] = instructions.ADDHL_BC

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.F & (1 << 6) != 0 {
        t.Error("N flag should be 0.")
    }

    if cpu.Registers.F & (1 << 5) == 0 {
        t.Error("H flag should be 1.")
    }

    if cpu.Registers.F & (1 << 4) == 0 {
        t.Error("C flag should be 1.")
    }

    if cpu.HL() != 0x124F {
        t.Error("HL should be 0x124F. Instead got: ", cpu.HL())
    }
}

func TestINC_BC(t *testing.T) {

    // Given
    cpu := InitSM83() 

    // When
    cpu.Registers.B = 0x25
    cpu.Registers.C = 0x26
    cpu.Memory.RAM[0x0100] = instructions.INC_BC

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.BC() != 0x2527 {
        t.Error("HL should be 0x2527. Instead got: ", cpu.BC())
    }
}

func TestDEC_BC(t *testing.T) {

    // Given
    cpu := InitSM83() 

    // When
    cpu.Registers.B = 0x25
    cpu.Registers.C = 0x26
    cpu.Memory.RAM[0x0100] = instructions.DEC_BC

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.BC() != 0x2525 {
        t.Error("HL should be 0x2525. Instead got: ", cpu.BC())
    }
}

func TestINC_DE(t *testing.T) {

    // Given
    cpu := InitSM83() 

    // When
    cpu.Registers.D = 0x25
    cpu.Registers.E = 0x26
    cpu.Memory.RAM[0x0100] = instructions.INC_DE

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.DE() != 0x2527 {
        t.Error("HL should be 0x2527. Instead got: ", cpu.DE())
    }
}

func TestDEC_DE(t *testing.T) {

    // Given
    cpu := InitSM83() 

    // When
    cpu.Registers.D = 0x25
    cpu.Registers.E = 0x26
    cpu.Memory.RAM[0x0100] = instructions.DEC_DE

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.DE() != 0x2525 {
        t.Error("HL should be 0x2525. Instead got: ", cpu.DE())
    }
}

func TestADDHL_DESetsHandCFlags(t *testing.T) {

    // Given
    cpu := InitSM83() 

    // When
    cpu.Registers.H = 0x23
    cpu.Registers.L = 0x20
    cpu.Registers.D = 0xEF
    cpu.Registers.E = 0x2F
    cpu.Memory.RAM[0x0100] = instructions.ADDHL_DE

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.F & (1 << 6) != 0 {
        t.Error("N flag should be 0.")
    }

    if cpu.Registers.F & (1 << 5) == 0 {
        t.Error("H flag should be 1.")
    }

    if cpu.Registers.F & (1 << 4) == 0 {
        t.Error("C flag should be 1.")
    }

    if cpu.HL() != 0x124F {
        t.Error("HL should be 0x124F. Instead got: ", cpu.HL())
    }
}
