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

func TestADDA_B(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.B = 0x35
    cpu.Memory.RAM[0x0100] = instructions.ADD_B

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if cpu.Registers.A != 0x6A {
        t.Error("A register should be 0x70. Instead got: ", cpu.Registers.A)
    }
}

func TestADDA_C(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.C = 0xCB
    cpu.Memory.RAM[0x0100] = instructions.ADD_C

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be set.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestADDA_D(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.D = 0x35
    cpu.Memory.RAM[0x0100] = instructions.ADD_D

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if cpu.Registers.A != 0x6A {
        t.Error("A register should be 0x70. Instead got: ", cpu.Registers.A)
    }
}

func TestADDA_E(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.E = 0x35
    cpu.Memory.RAM[0x0100] = instructions.ADD_E

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if cpu.Registers.A != 0x6A {
        t.Error("A register should be 0x70. Instead got: ", cpu.Registers.A)
    }
}

func TestADDA_H(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.H = 0x35
    cpu.Memory.RAM[0x0100] = instructions.ADD_H

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if cpu.Registers.A != 0x6A {
        t.Error("A register should be 0x70. Instead got: ", cpu.Registers.A)
    }
}

func TestADDA_L(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.L = 0x35
    cpu.Memory.RAM[0x0100] = instructions.ADD_L

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if cpu.Registers.A != 0x6A {
        t.Error("A register should be 0x70. Instead got: ", cpu.Registers.A)
    }
}

func TestADDA_A(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Memory.RAM[0x0100] = instructions.ADD_A

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if cpu.Registers.A != 0x6A {
        t.Error("A register should be 0x70. Instead got: ", cpu.Registers.A)
    }
}

func TestADDA_HL(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.H = 0x90
    cpu.Registers.L = 0x08
    cpu.Memory.RAM[0x0100] = instructions.ADD_indHL
    cpu.Memory.RAM[0x9008] = 0x35

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if cpu.Registers.A != 0x6A {
        t.Error("A register should be 0x70. Instead got: ", cpu.Registers.A)
    }
}

func TestADC_B(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.B = 0xCA
    cpu.Memory.RAM[0x0100] = instructions.CCF
    // flag C is set
    cpu.Memory.RAM[0x0101] = instructions.ADC_B

    expectedCycles := 1 + 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 4)) == 0 {
        t.Error("C flag should be 1.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestADC_C(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.C = 0xCA
    cpu.Memory.RAM[0x0100] = instructions.CCF
    // flag C is set
    cpu.Memory.RAM[0x0101] = instructions.ADC_C

    expectedCycles := 1 + 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 4)) == 0 {
        t.Error("C flag should be 1.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestADC_D(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.D = 0xCA
    cpu.Memory.RAM[0x0100] = instructions.CCF
    // flag C is set
    cpu.Memory.RAM[0x0101] = instructions.ADC_D

    expectedCycles := 1 + 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 4)) == 0 {
        t.Error("C flag should be 1.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestADC_E(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.E = 0xCA
    cpu.Memory.RAM[0x0100] = instructions.CCF
    // flag C is set
    cpu.Memory.RAM[0x0101] = instructions.ADC_E

    expectedCycles := 1 + 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 4)) == 0 {
        t.Error("C flag should be 1.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestADC_H(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.H = 0xCA
    cpu.Memory.RAM[0x0100] = instructions.CCF
    // flag C is set
    cpu.Memory.RAM[0x0101] = instructions.ADC_H

    expectedCycles := 1 + 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 4)) == 0 {
        t.Error("C flag should be 1.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestADC_L(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.L = 0xCA
    cpu.Memory.RAM[0x0100] = instructions.CCF
    // flag C is set
    cpu.Memory.RAM[0x0101] = instructions.ADC_L

    expectedCycles := 1 + 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 4)) == 0 {
        t.Error("C flag should be 1.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestADC_A(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x80
    cpu.Memory.RAM[0x0100] = instructions.CCF
    // flag C is set
    cpu.Memory.RAM[0x0101] = instructions.ADC_A

    expectedCycles := 1 + 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) == 0 {
        t.Error("C flag should be 1.")
    }

    if cpu.Registers.A != 0x01 {
        t.Error("A register should be 1. Instead got: ", cpu.Registers.A)
    }
}

func TestADC_HL(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.H = 0x90
    cpu.Registers.L = 0x08
    cpu.Memory.RAM[0x0100] = instructions.CCF
    cpu.Memory.RAM[0x0101] = instructions.ADC_indHL
    cpu.Memory.RAM[0x9008] = 0x35

    expectedCycles := 1 + 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if cpu.Registers.A != 0x6B {
        t.Error("A register should be 0x6B. Instead got: ", cpu.Registers.A)
    }
}

func TestSUB_B(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.B = 0x35
    cpu.Memory.RAM[0x0100] = instructions.SUB_B

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}


// TestSUB_BSetsHandCflags verifies that Half and Carry flag are set.
// Do this once for SUB_B. Function that subtracts is common to all SUB operations.
func TestSUB_BSetsHandCflags(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.B = 0x47
    cpu.Memory.RAM[0x0100] = instructions.SUB_B

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 4)) == 0 {
        t.Error("C flag should be 1.")
    }

    if cpu.Registers.A != 0xEE {
        t.Error("A register should be 0xEE. Instead got: ", cpu.Registers.A)
    }
}

func TestSUB_C(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.C = 0x35
    cpu.Memory.RAM[0x0100] = instructions.SUB_C

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestSUB_D(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.D = 0x35
    cpu.Memory.RAM[0x0100] = instructions.SUB_D

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestSUB_E(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.E = 0x35
    cpu.Memory.RAM[0x0100] = instructions.SUB_E

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestSUB_H(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.H = 0x35
    cpu.Memory.RAM[0x0100] = instructions.SUB_H

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestSUB_L(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.L = 0x35
    cpu.Memory.RAM[0x0100] = instructions.SUB_L

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestSUB_A(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Memory.RAM[0x0100] = instructions.SUB_A

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestSUB_indHL(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.H = 0x90
    cpu.Registers.L = 0x08
    cpu.Memory.RAM[0x0100] = instructions.SUB_indHL
    cpu.Memory.RAM[0x9008] = 0x35

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}


func TestSBC_B(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.B = 0x34
    cpu.Memory.RAM[0x0100] = instructions.CCF
    // flag C is set
    cpu.Memory.RAM[0x0101] = instructions.SBC_B

    expectedCycles := 1 + 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestSBC_BSetsHandCflags(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.B = 0x46
    cpu.Memory.RAM[0x0100] = instructions.CCF
    // flag C is set
    cpu.Memory.RAM[0x0101] = instructions.SBC_B

    expectedCycles := 1 + 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 4)) == 0 {
        t.Error("C flag should be 1.")
    }

    if cpu.Registers.A != 0xEE {
        t.Error("A register should be 0xEE. Instead got: ", cpu.Registers.A)
    }
}

func TestSBC_C(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.C = 0x34
    cpu.Memory.RAM[0x0100] = instructions.CCF
    // flag C is set
    cpu.Memory.RAM[0x0101] = instructions.SBC_C

    expectedCycles := 1 + 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestSBC_D(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.D = 0x34
    cpu.Memory.RAM[0x0100] = instructions.CCF
    // flag C is set
    cpu.Memory.RAM[0x0101] = instructions.SBC_D

    expectedCycles := 1 + 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestSBC_E(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.E = 0x34
    cpu.Memory.RAM[0x0100] = instructions.CCF
    // flag C is set
    cpu.Memory.RAM[0x0101] = instructions.SBC_E

    expectedCycles := 1 + 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestSBC_A(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    // flag C is set
    cpu.Memory.RAM[0x0100] = instructions.SBC_A

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestSBC_H(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.H = 0x34
    cpu.Memory.RAM[0x0100] = instructions.CCF
    // flag C is set
    cpu.Memory.RAM[0x0101] = instructions.SBC_H

    expectedCycles := 1 + 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestSBC_L(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.L = 0x34
    cpu.Memory.RAM[0x0100] = instructions.CCF
    // flag C is set
    cpu.Memory.RAM[0x0101] = instructions.SBC_L

    expectedCycles := 1 + 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestSBC_indHL(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.H = 0x90
    cpu.Registers.L = 0x08
    cpu.Memory.RAM[0x0100] = instructions.SBC_indHL
    cpu.Memory.RAM[0x9008] = 0x35

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestAND_B(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.B = 0x13
    cpu.Memory.RAM[0x0100] = instructions.AND_B

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if cpu.Registers.A != 0x11 {
        t.Error("A register should be 0x11. Instead got: ", cpu.Registers.A)
    }
}

func TestAND_C(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.C = 0x13
    cpu.Memory.RAM[0x0100] = instructions.AND_C

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if cpu.Registers.A != 0x11 {
        t.Error("A register should be 0x11. Instead got: ", cpu.Registers.A)
    }
}

func TestAND_D(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.D = 0x13
    cpu.Memory.RAM[0x0100] = instructions.AND_D

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if cpu.Registers.A != 0x11 {
        t.Error("A register should be 0x11. Instead got: ", cpu.Registers.A)
    }
}

func TestAND_E(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.E = 0x13
    cpu.Memory.RAM[0x0100] = instructions.AND_E

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if cpu.Registers.A != 0x11 {
        t.Error("A register should be 0x11. Instead got: ", cpu.Registers.A)
    }
}

func TestAND_H(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.H = 0x13
    cpu.Memory.RAM[0x0100] = instructions.AND_H

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if cpu.Registers.A != 0x11 {
        t.Error("A register should be 0x11. Instead got: ", cpu.Registers.A)
    }
}

func TestAND_L(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.L = 0x13
    cpu.Memory.RAM[0x0100] = instructions.AND_L

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if cpu.Registers.A != 0x11 {
        t.Error("A register should be 0x11. Instead got: ", cpu.Registers.A)
    }
}

func TestAND_A(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Memory.RAM[0x0100] = instructions.AND_A

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if cpu.Registers.A != 0x35 {
        t.Error("A register should be 0x35. Instead got: ", cpu.Registers.A)
    }
}

func TestAND_indHL(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.H = 0x99
    cpu.Registers.L = 0x13
    cpu.Memory.RAM[0x0100] = instructions.AND_indHL
    cpu.Memory.RAM[0x9913] = 0x13

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if cpu.Registers.A != 0x11 {
        t.Error("A register should be 0x11. Instead got: ", cpu.Registers.A)
    }
}

func TestAND_AWorksWithZeroValue(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x0
    cpu.Memory.RAM[0x0100] = instructions.AND_A

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestXOR_B(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.B = 0x13
    cpu.Memory.RAM[0x0100] = instructions.XOR_B

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if cpu.Registers.A != 0x26 {
        t.Error("A register should be 0x26. Instead got: ", cpu.Registers.A)
    }
}

func TestXOR_C(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.C = 0x13
    cpu.Memory.RAM[0x0100] = instructions.XOR_C

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if cpu.Registers.A != 0x26 {
        t.Error("A register should be 0x26. Instead got: ", cpu.Registers.A)
    }
}

func TestXOR_D(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.D = 0x13
    cpu.Memory.RAM[0x0100] = instructions.XOR_D

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if cpu.Registers.A != 0x26 {
        t.Error("A register should be 0x26. Instead got: ", cpu.Registers.A)
    }
}

func TestXOR_E(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.E = 0x13
    cpu.Memory.RAM[0x0100] = instructions.XOR_E

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if cpu.Registers.A != 0x26 {
        t.Error("A register should be 0x26. Instead got: ", cpu.Registers.A)
    }
}

func TestXOR_H(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.H = 0x13
    cpu.Memory.RAM[0x0100] = instructions.XOR_H

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if cpu.Registers.A != 0x26 {
        t.Error("A register should be 0x26. Instead got: ", cpu.Registers.A)
    }
}

func TestXOR_L(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.L = 0x13
    cpu.Memory.RAM[0x0100] = instructions.XOR_L

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if cpu.Registers.A != 0x26 {
        t.Error("A register should be 0x26. Instead got: ", cpu.Registers.A)
    }
}

func TestXOR_A(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Memory.RAM[0x0100] = instructions.XOR_A

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestXOR_indHL(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.H = 0x99
    cpu.Registers.L = 0x13
    cpu.Memory.RAM[0x0100] = instructions.XOR_indHL
    cpu.Memory.RAM[0x9913] = 0x13

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if cpu.Registers.A != 0x26 {
        t.Error("A register should be 0x11. Instead got: ", cpu.Registers.A)
    }
}

func TestOR_B(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.B = 0x13
    cpu.Memory.RAM[0x0100] = instructions.OR_B

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if cpu.Registers.A != 0x37 {
        t.Error("A register should be 0x37. Instead got: ", cpu.Registers.A)
    }
}

func TestOR_C(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.C = 0x13
    cpu.Memory.RAM[0x0100] = instructions.OR_C

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if cpu.Registers.A != 0x37 {
        t.Error("A register should be 0x37. Instead got: ", cpu.Registers.A)
    }
}


func TestOR_D(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.D = 0x13
    cpu.Memory.RAM[0x0100] = instructions.OR_D

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if cpu.Registers.A != 0x37 {
        t.Error("A register should be 0x37. Instead got: ", cpu.Registers.A)
    }
}


func TestOR_E(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.E = 0x13
    cpu.Memory.RAM[0x0100] = instructions.OR_E

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if cpu.Registers.A != 0x37 {
        t.Error("A register should be 0x37. Instead got: ", cpu.Registers.A)
    }
}


func TestOR_H(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.H = 0x13
    cpu.Memory.RAM[0x0100] = instructions.OR_H

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if cpu.Registers.A != 0x37 {
        t.Error("A register should be 0x37. Instead got: ", cpu.Registers.A)
    }
}

func TestOR_L(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.L = 0x13
    cpu.Memory.RAM[0x0100] = instructions.OR_L

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if cpu.Registers.A != 0x37 {
        t.Error("A register should be 0x37. Instead got: ", cpu.Registers.A)
    }
}

func TestOR_indHL(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.H = 0x99
    cpu.Registers.L = 0x13
    cpu.Memory.RAM[0x0100] = instructions.OR_indHL
    cpu.Memory.RAM[0x9913] = 0x13

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if cpu.Registers.A != 0x37 {
        t.Error("A register should be 0x37. Instead got: ", cpu.Registers.A)
    }
}

func TestOR_A(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Memory.RAM[0x0100] = instructions.OR_A

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if cpu.Registers.A != 0x35 {
        t.Error("A register should be 0x35. Instead got: ", cpu.Registers.A)
    }
}

func TestCP_B(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.B = 0x35
    cpu.Memory.RAM[0x0100] = instructions.CP_A

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if cpu.Registers.A != 0x35 {
        t.Error("A register should be 0x35. Instead got: ", cpu.Registers.A)
    }
}


// TestSUB_BSetsHandCflags verifies that Half and Carry flag are set.
// Do this once for SUB_B. Function that subtracts is common to all SUB operations.
func TestCP_BSetsHandCflags(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.B = 0x47
    cpu.Memory.RAM[0x0100] = instructions.CP_B

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 4)) == 0 {
        t.Error("C flag should be 1.")
    }
}

func TestCP_C(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.C = 0x35
    cpu.Memory.RAM[0x0100] = instructions.CP_C

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }
}

func TestCP_D(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.D = 0x35
    cpu.Memory.RAM[0x0100] = instructions.CP_D

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }
}

func TestCP_E(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.E = 0x35
    cpu.Memory.RAM[0x0100] = instructions.CP_E

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }
}

func TestCP_H(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.H = 0x35
    cpu.Memory.RAM[0x0100] = instructions.CP_H

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }
}

func TestCP_L(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.L = 0x35
    cpu.Memory.RAM[0x0100] = instructions.CP_L

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }
}

func TestCP_A(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Memory.RAM[0x0100] = instructions.CP_A

    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }
}

func TestCP_indHL(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Registers.H = 0x90
    cpu.Registers.L = 0x08
    cpu.Memory.RAM[0x0100] = instructions.CP_indHL
    cpu.Memory.RAM[0x9008] = 0x35

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }
}

func TestADD_d8(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Memory.RAM[0x0100] = instructions.ADD_d8
    cpu.Memory.RAM[0x0101] = 0x35

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if cpu.Registers.A != 0x6A {
        t.Error("A register should be 0x70. Instead got: ", cpu.Registers.A)
    }
}

func TestADC_d8(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Memory.RAM[0x0100] = instructions.CCF
    // flag C is set
    cpu.Memory.RAM[0x0101] = instructions.ADC_d8
    cpu.Memory.RAM[0x0102] = 0xCA

    expectedCycles := 1 + 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 4)) == 0 {
        t.Error("C flag should be 1.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestSUB_d8(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Memory.RAM[0x0100] = instructions.SUB_d8
    cpu.Memory.RAM[0x0101] = 0x35

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestSBC_d8(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Memory.RAM[0x0100] = instructions.CCF
    // flag C is set
    cpu.Memory.RAM[0x0101] = instructions.SBC_d8
    cpu.Memory.RAM[0x0102] = 0x34

    expectedCycles := 1 + 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if cpu.Registers.A != 0x00 {
        t.Error("A register should be 0. Instead got: ", cpu.Registers.A)
    }
}

func TestAND_d8(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Memory.RAM[0x0100] = instructions.AND_d8
    cpu.Memory.RAM[0x0101] = 0x13

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) == 0 {
        t.Error("H flag should be 1.")
    }

    if cpu.Registers.A != 0x11 {
        t.Error("A register should be 0x11. Instead got: ", cpu.Registers.A)
    }
}

func TestXOR_d8(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Memory.RAM[0x0100] = instructions.XOR_d8
    cpu.Memory.RAM[0x0101] = 0x13

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if cpu.Registers.A != 0x26 {
        t.Error("A register should be 0x26. Instead got: ", cpu.Registers.A)
    }
}

func TestOR_d8(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Memory.RAM[0x0100] = instructions.OR_d8
    cpu.Memory.RAM[0x0101] = 0x13

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 7)) != 0 {
        t.Error("Z flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 6)) != 0 {
        t.Error("N flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if cpu.Registers.A != 0x37 {
        t.Error("A register should be 0x37. Instead got: ", cpu.Registers.A)
    }
}

func TestCP_d8(t *testing.T) {

    // Given
    cpu := InitSM83()

    // When
    cpu.Registers.A = 0x35
    cpu.Memory.RAM[0x0100] = instructions.CP_d8
    cpu.Memory.RAM[0x0101] = 0x35

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if (cpu.Registers.F & (1 << 6)) == 0 {
        t.Error("N flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 7)) == 0 {
        t.Error("Z flag should be 1.")
    }

    if (cpu.Registers.F & (1 << 5)) != 0 {
        t.Error("H flag should be 0.")
    }

    if (cpu.Registers.F & (1 << 4)) != 0 {
        t.Error("C flag should be 0.")
    }

    if cpu.Registers.A != 0x35 {
        t.Error("A register should be 0x35. Instead got: ", cpu.Registers.A)
    }
}
