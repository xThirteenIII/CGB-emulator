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

func TestINC_A(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x34
    cpu.Memory.RAM[0x0100] = instructions.INC_A

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.A != 0x35 {
        t.Error("A register should be 0x35, instead got: ", cpu.Registers.A)
    }
}

func TestINC_ASetsZandHFlags(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0xFF
    cpu.Memory.RAM[0x0100] = instructions.INC_A

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0, instead got: ", cpu.Registers.A)
    }

    if cpu.Registers.F != 0b10100000 {
        t.Error("Z and H should be set. Instead got: ", cpu.Registers.F)
    }
}

func TestDEC_A(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x34
    cpu.Memory.RAM[0x0100] = instructions.DEC_A

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.A != 0x33 {
        t.Error("A register should be 0x33, instead got: ", cpu.Registers.A)
    }
}

func TestDEC_ASetsZandHFlags(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0xF0
    cpu.Memory.RAM[0x0100] = instructions.DEC_A

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.A != 0xEF {
        t.Error("A register should be 0xEF, instead got: ", cpu.Registers.A)
    }

    if cpu.Registers.F != 0b01100000 {
        t.Error("Z, N and H should be set. Instead got: ", cpu.Registers.F)
    }
}

func TestDAAAdjustsAferINC_A(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Memory.RAM[0x0100] = instructions.LDA_d8
    cpu.Memory.RAM[0x0101] = 0x09
    cpu.Memory.RAM[0x0102] = instructions.INC_A
    cpu.Memory.RAM[0x0103] = instructions.DAA

    expectedCycles := 2 + 1 + 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.A != 0x10 {
        t.Error("A register should be 0x10. Instead got: ", cpu.Registers.A)
    }
}

func TestDAAAdjustsAferDEC_A(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Memory.RAM[0x0100] = instructions.LDA_d8
    cpu.Memory.RAM[0x0101] = 0x20
    cpu.Memory.RAM[0x0102] = instructions.DEC_A
    cpu.Memory.RAM[0x0103] = instructions.DAA

    expectedCycles := 2 + 1 + 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    // A = 0x1F = 0x19
    if cpu.Registers.A != 0x19 {
        t.Error("A register should be 0x19. Instead got: ", cpu.Registers.A)
    }
}

func TestCPL(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x3F
    cpu.Memory.RAM[0x0100] = instructions.CPL

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    // A = 0x3F = 0x00111111
    // ~A = 0x11000000
    if cpu.Registers.A != 0xC0 {
        t.Error("A register should be 0xC0. Instead got: ", cpu.Registers.A)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N register should be set.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H register should be set.")
    }
}

func TestINC_indHL(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.H = 0x30
    cpu.Registers.L = 0x59
    cpu.Memory.RAM[0x0100] = instructions.INC_indHL
    cpu.Memory.RAM[0x3059] = 0x69

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Memory.RAM[0x3059] != 0x6A {
        t.Error("Data at 0x3059 should be 0x6A. Instead got: ", cpu.Memory.RAM[0x3059])
    }
}

func TestINC_indHLSetsHAndZFlags(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.H = 0x30
    cpu.Registers.L = 0x59
    cpu.Memory.RAM[0x0100] = instructions.INC_indHL
    cpu.Memory.RAM[0x3059] = 0xFF

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Memory.RAM[0x3059] != 0x00 {
        t.Error("Data at 0x3059 should be 0x00. Instead got: ", cpu.Memory.RAM[0x3059])
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be set")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be set.")
    }
}

func TestDEC_indHL(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.H = 0x30
    cpu.Registers.L = 0x59
    cpu.Memory.RAM[0x0100] = instructions.DEC_indHL
    cpu.Memory.RAM[0x3059] = 0x69

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Memory.RAM[0x3059] != 0x68 {
        t.Error("Data at 0x3059 should be 0x68. Instead got: ", cpu.Memory.RAM[0x3059])
    }
}

func TestDEC_indHLSetsZAndNFlags(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.H = 0x30
    cpu.Registers.L = 0x59
    cpu.Memory.RAM[0x0100] = instructions.DEC_indHL
    cpu.Memory.RAM[0x3059] = 0x01

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be set")
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be set.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should not be set.")
    }
}

func TestSCF(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.SetNflag()
    cpu.SetHflag()
    cpu.ClearCflag()
    cpu.Memory.RAM[0x0100] = instructions.SCF

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) == 0 {
        t.Error("C flag should be set.")
    }
}

func TestCCF(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.SetNflag()
    cpu.SetHflag()
    cpu.ClearCflag()
    cpu.Memory.RAM[0x0100] = instructions.SCF

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) == 0 {
        t.Error("C flag should be set.")
    }
}
