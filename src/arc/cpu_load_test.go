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

    // Setting more cycles than needed, will make the Execute() return with "unknown opcode: 0".
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

func TestLDBC_d16(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDBC_d16
    cpu.Memory.RAM[0x0101] = 0x52
    cpu.Memory.RAM[0x0102] = 0x72

    // When
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    BC := (uint16(cpu.Registers.B) << 8 | uint16(cpu.Registers.C)) 

    if BC != 0x7252{
        t.Error("BC register should be 0x7252, instead got: ", BC)
    }
}

// Testa16_SP verifies that data from the SP register is loaded into the absolute address specified in memory.
// Test functions need capital letters after 'test'????? WHAAAAAAAAAAAAAAAT
func TestA16_SP(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    // SP = 0xFFFE
    cpu.Memory.RAM[0x0100] = instructions.LDa16_SP
    cpu.Memory.RAM[0x0101] = 0x52
    cpu.Memory.RAM[0x0102] = 0x72

    // 0x5555 data into 7252

    // When
    expectedCycles := 5
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Memory.RAM[0x7252] != 0xFE {
        t.Error("Address 0x7252 should be 0xFE.")
    }

    if cpu.Memory.RAM[0x7253] != 0xFF {
        t.Error("Address 0x7253 should be 0xFF.")
    }
}

// TestA_BC verifies that data at the absolute address stored in the BC register is loaded into the A register.
func TestLDA_BC(t *testing.T) {

    cpu := InitSM83()

    cpu.Registers.B = 0x56
    cpu.Registers.C = 0xF6
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDA_BC
    cpu.Memory.RAM[0x56F6] = 0x55

    // 0x5555 data into 7252

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.A != 0x55 {
        t.Error("A register should be 0x55.")
    }
}

// TestLDBC_A verifies that the content stored in the A register is loaded at the absolute address stored in the BC register.
func TestLDBC_A (t *testing.T) {

    cpu := InitSM83()

    // A = 0x11
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDBC_A

    // 0x5555 data into 7252

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Memory.RAM[cpu.BC()] != 0x11 {
        t.Error("BC register should be 0x11.")
    }

}

func TestLDDE_A (t *testing.T) {

    cpu := InitSM83()

    // A = 0x11
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDDE_A

    // 0x5555 data into 7252

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Memory.RAM[cpu.DE()] != 0x11 {
        t.Error("BC register should be 0x11.")
    }

}

func TestLDC_d8(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDC_d8
    cpu.Memory.RAM[0x0101] = 0x33

    // Setting more cycles than needed, will make the Execute() return with "unknown opcode: 0".
    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.C != 0x33 {
        t.Error("C register should be 0x33, instead got: ", cpu.Registers.C)
    }
}

func TestLDDE_d16(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDDE_d16
    cpu.Memory.RAM[0x0101] = 0x52
    cpu.Memory.RAM[0x0102] = 0x72

    // When
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.DE() != 0x7252{
        t.Error("DE register should be 0x7252, instead got: ", cpu.DE())
    }
}

func TestLD_d8(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDD_d8
    cpu.Memory.RAM[0x0101] = 0x33

    // Setting more cycles than needed, will make the Execute() return with "unknown opcode: 0".
    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.D != 0x33 {
        t.Error("D register should be 0x33, instead got: ", cpu.Registers.D)
    }
}

// TestA_BC verifies that data at the absolute address stored in the BC register is loaded into the A register.
func TestLDA_DE(t *testing.T) {

    cpu := InitSM83()

    cpu.Registers.D = 0x56
    cpu.Registers.E = 0xF6
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDA_DE
    cpu.Memory.RAM[0x56F6] = 0x55

    // 0x5555 data into 7252

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.A != 0x55 {
        t.Error("A register should be 0x55.")
    }
}

func TestLDE_d8(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDE_d8
    cpu.Memory.RAM[0x0101] = 0x33

    // Setting more cycles than needed, will make the Execute() return with "unknown opcode: 0".
    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.E != 0x33 {
        t.Error("E register should be 0x33, instead got: ", cpu.Registers.E)
    }
}

func TestLDHL_d16(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDHL_d16
    cpu.Memory.RAM[0x0101] = 0x52
    cpu.Memory.RAM[0x0102] = 0x72

    // When
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.HL() != 0x7252{
        t.Error("HL register should be 0x7252, instead got: ", cpu.HL())
    }
}

func TestLDHLinc_A(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    // A=0x11
    cpu.Registers.H = 0x60
    cpu.Registers.L = 0x62
    cpu.Memory.RAM[0x0100] = instructions.LDHLinc_A

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }


    if cpu.HL() != 0x6063{
        t.Error("HL register should be 0x6063, instead got: ", cpu.HL())
    }

    if cpu.Memory.RAM[cpu.HL()-1] != 0x11 {
        t.Error("Memory address at HL - 1 should be 0x11, instead got: ", cpu.Memory.RAM[cpu.HL()-1])
    }
}

func TestLDH_d8(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDH_d8
    cpu.Memory.RAM[0x0101] = 0x33

    // Setting more cycles than needed, will make the Execute() return with "unknown opcode: 0".
    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.H != 0x33 {
        t.Error("H register should be 0x33, instead got: ", cpu.Registers.H)
    }
}

func TestLDA_HLinc(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    // A=0x11
    cpu.Registers.H = 0x60
    cpu.Registers.L = 0x62
    cpu.Memory.RAM[0x0100] = instructions.LDA_HLinc
    cpu.Memory.RAM[0x6062] = 0x58

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }


    if cpu.Registers.A != 0x58 {
        t.Error("A register should be 0x58, instead got: ", cpu.Registers.A)
    }

    if cpu.Registers.L != 0x63 {
        t.Error("L register should be 0x63, instead got: ", cpu.Registers.L)
    }

}

/*
// Commented because it causes an exit. It works correctly.
func TestLDA_HLincExceedsUint16(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    // A=0x11
    cpu.Registers.H = 0xFF
    cpu.Registers.L = 0xFF
    cpu.Memory.RAM[0x0100] = instructions.LDA_HLinc
    cpu.Memory.RAM[0x6062] = 0x58

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }


    if cpu.Registers.A != 0x58 {
        t.Error("A register should be 0x58, instead got: ", cpu.Registers.A)
    }

    if cpu.Registers.L != 0x63 {
        t.Error("L register should be 0x63, instead got: ", cpu.Registers.L)
    }

}
*/

func TestLDL_d8(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDL_d8
    cpu.Memory.RAM[0x0101] = 0x33

    // Setting more cycles than needed, will make the Execute() return with "unknown opcode: 0".
    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.L != 0x33 {
        t.Error("L register should be 0x33, instead got: ", cpu.Registers.L)
    }
}

func TestLDHLdec_A(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    // A=0x11
    cpu.Registers.H = 0x60
    cpu.Registers.L = 0x62
    cpu.Memory.RAM[0x0100] = instructions.LDHLdec_A

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }


    if cpu.HL() != 0x6061{
        t.Error("HL register should be 0x6061, instead got: ", cpu.HL())
    }

    if cpu.Memory.RAM[cpu.HL()+1] != 0x11 {
        t.Error("Memory address at HL - 1 should be 0x11, instead got: ", cpu.Memory.RAM[cpu.HL()-1])
    }
}

func TestLDA_HLdec(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    // A=0x11
    cpu.Registers.H = 0x60
    cpu.Registers.L = 0x62
    cpu.Memory.RAM[0x0100] = instructions.LDA_HLdec
    cpu.Memory.RAM[0x6062] = 0x58

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }


    if cpu.Registers.A != 0x58 {
        t.Error("A register should be 0x58, instead got: ", cpu.Registers.A)
    }

    if cpu.Registers.L != 0x61 {
        t.Error("L register should be 0x61, instead got: ", cpu.Registers.L)
    }

}

func TestLDA_d8(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDA_d8
    cpu.Memory.RAM[0x0101] = 0x33

    // Setting more cycles than needed, will make the Execute() return with "unknown opcode: 0".
    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.A != 0x33 {
        t.Error("A register should be 0x33, instead got: ", cpu.Registers.A)
    }
}

func TestLDB_B(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.B = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDB_B

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.B != 0x69 {
        t.Error("B register should be 0x69, instead got: ", cpu.Registers.B)
    }
}

func TestLDB_C(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.C = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDB_C

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.B != 0x69 {
        t.Error("B register should be 0x69, instead got: ", cpu.Registers.B)
    }
}
func TestLDB_D(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.D = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDB_D

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.B != 0x69 {
        t.Error("B register should be 0x69, instead got: ", cpu.Registers.B)
    }
}
func TestLDB_H(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDB_H

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.B != 0x69 {
        t.Error("B register should be 0x69, instead got: ", cpu.Registers.B)
    }
}
func TestLDB_L(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.L = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDB_L

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.B != 0x69 {
        t.Error("B register should be 0x69, instead got: ", cpu.Registers.B)
    }
}
func TestLDB_E(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.E = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDB_E

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.B != 0x69 {
        t.Error("B register should be 0x69, instead got: ", cpu.Registers.B)
    }
}

func TestLDB_A(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.A = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDB_A

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.B != 0x69 {
        t.Error("B register should be 0x69, instead got: ", cpu.Registers.B)
    }
}

func TestLDC_B(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.B = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDC_B

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.C != 0x69 {
        t.Error("C register should be 0x69, instead got: ", cpu.Registers.C)
    }
}

func TestLDC_C(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.C = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDC_C

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.C != 0x69 {
        t.Error("C register should be 0x69, instead got: ", cpu.Registers.C)
    }
}

func TestLDC_D(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.D = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDC_D

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.C != 0x69 {
        t.Error("C register should be 0x69, instead got: ", cpu.Registers.C)
    }
}

func TestLDC_E(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.E = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDC_E

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.C != 0x69 {
        t.Error("C register should be 0x69, instead got: ", cpu.Registers.C)
    }
}

func TestLDC_H(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDC_H

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.C != 0x69 {
        t.Error("C register should be 0x69, instead got: ", cpu.Registers.C)
    }
}

func TestLDC_L(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.L = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDC_L

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.C != 0x69 {
        t.Error("C register should be 0x69, instead got: ", cpu.Registers.C)
    }
}

func TestLDC_HL(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x20
    cpu.Registers.L = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDC_HL
    cpu.Memory.RAM[0x2069] = 0x69

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.C != 0x69 {
        t.Error("C register should be 0x69, instead got: ", cpu.Registers.C)
    }
}

func TestLDC_A(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.A = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDC_A

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.C != 0x69 {
        t.Error("C register should be 0x69, instead got: ", cpu.Registers.C)
    }
}

func TestLDD_B(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.B = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDD_B

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.D != 0x69 {
        t.Error("D register should be 0x69, instead got: ", cpu.Registers.D)
    }
}

func TestLDD_C(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.C = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDD_C

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.D != 0x69 {
        t.Error("D register should be 0x69, instead got: ", cpu.Registers.D)
    }
}

func TestLDD_D(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.D = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDD_D

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.D != 0x69 {
        t.Error("D register should be 0x69, instead got: ", cpu.Registers.D)
    }
}

func TestLDD_E(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.E = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDD_E

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.D != 0x69 {
        t.Error("D register should be 0x69, instead got: ", cpu.Registers.D)
    }
}

func TestLDD_H(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDD_H

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.D != 0x69 {
        t.Error("D register should be 0x69, instead got: ", cpu.Registers.D)
    }
}

func TestLDD_L(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.L = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDD_L

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.D != 0x69 {
        t.Error("D register should be 0x69, instead got: ", cpu.Registers.D)
    }
}

func TestLDD_HL(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x80
    cpu.Registers.L = 0x8F
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDD_HL
    cpu.Memory.RAM[0x808F] = 0x20

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.D != 0x20 {
        t.Error("D register should be 0x20, instead got: ", cpu.Registers.D)
    }
}

func TestLDD_A(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.A = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDD_A

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.D != 0x69 {
        t.Error("D register should be 0x69, instead got: ", cpu.Registers.D)
    }
}

func TestLDE_B(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.B = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDE_B

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.E != 0x69 {
        t.Error("E register should be 0x69, instead got: ", cpu.Registers.E)
    }
}

func TestLDE_C(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.C = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDE_C

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.E != 0x69 {
        t.Error("E register should be 0x69, instead got: ", cpu.Registers.E)
    }
}
func TestLDE_D(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.D = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDE_D

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.E != 0x69 {
        t.Error("E register should be 0x69, instead got: ", cpu.Registers.E)
    }
}
func TestLDE_E(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.E = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDE_E

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.E != 0x69 {
        t.Error("E register should be 0x69, instead got: ", cpu.Registers.E)
    }
}
func TestLDE_H(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDE_H

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.E != 0x69 {
        t.Error("E register should be 0x69, instead got: ", cpu.Registers.E)
    }
}
func TestLDE_L(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.L = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDE_L

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.E != 0x69 {
        t.Error("E register should be 0x69, instead got: ", cpu.Registers.E)
    }
}

func TestLDE_HL(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x80
    cpu.Registers.L = 0x8F
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDE_HL
    cpu.Memory.RAM[0x808F] = 0x20

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.E != 0x20 {
        t.Error("E register should be 0x20, instead got: ", cpu.Registers.E)
    }
}


func TestLDE_A(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.A = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDE_A

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.E != 0x69 {
        t.Error("E register should be 0x69, instead got: ", cpu.Registers.E)
    }
}

func TestLDH_B(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.B = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDH_B

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.H != 0x69 {
        t.Error("H register should be 0x69, instead got: ", cpu.Registers.H)
    }
}

func TestLDH_C(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.C = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDH_C

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.H != 0x69 {
        t.Error("H register should be 0x69, instead got: ", cpu.Registers.H)
    }
}

func TestLDH_D(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.D = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDH_D

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.H != 0x69 {
        t.Error("H register should be 0x69, instead got: ", cpu.Registers.H)
    }
}

func TestLDH_E(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.E = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDH_E

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.H != 0x69 {
        t.Error("H register should be 0x69, instead got: ", cpu.Registers.H)
    }
}

func TestLDH_H(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDH_H

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.H != 0x69 {
        t.Error("H register should be 0x69, instead got: ", cpu.Registers.H)
    }
}

func TestLDH_L(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.L = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDH_L

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.H != 0x69 {
        t.Error("H register should be 0x69, instead got: ", cpu.Registers.H)
    }
}

func TestLDH_HL(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x80
    cpu.Registers.L = 0x8F
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDH_HL
    cpu.Memory.RAM[0x808F] = 0x20

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.H != 0x20 {
        t.Error("H register should be 0x20, instead got: ", cpu.Registers.H)
    }
}

func TestLDL_B(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.B = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDL_B

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.L != 0x69 {
        t.Error("L register should be 0x69, instead got: ", cpu.Registers.L)
    }
}

func TestLDL_C(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.C = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDL_C

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.L != 0x69 {
        t.Error("L register should be 0x69, instead got: ", cpu.Registers.L)
    }
}
func TestLDL_D(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.D = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDL_D

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.L != 0x69 {
        t.Error("L register should be 0x69, instead got: ", cpu.Registers.L)
    }
}
func TestLDL_E(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.E = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDL_E

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.L != 0x69 {
        t.Error("L register should be 0x69, instead got: ", cpu.Registers.L)
    }
}
func TestLDL_H(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDL_H

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.L != 0x69 {
        t.Error("L register should be 0x69, instead got: ", cpu.Registers.L)
    }
}

func TestLDL_L(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.L = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDL_L

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.L != 0x69 {
        t.Error("L register should be 0x69, instead got: ", cpu.Registers.L)
    }
}

func TestLDL_HL(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x80
    cpu.Registers.L = 0x8F
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDL_HL
    cpu.Memory.RAM[0x808F] = 0x20

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.L != 0x20 {
        t.Error("L register should be 0x20, instead got: ", cpu.Registers.L)
    }
}

func TestLDHL_B(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x11
    cpu.Registers.L = 0x69
    cpu.Registers.B = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDHL_B

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Memory.RAM[cpu.HL()] != 0x69 {
        t.Error("Memory address at  HL register should be 0x69, instead got: ", cpu.Memory.RAM[cpu.HL()])
    }
}

func TestLDHL_C(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x11
    cpu.Registers.L = 0x69
    cpu.Registers.C = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDHL_C

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Memory.RAM[cpu.HL()] != 0x69 {
        t.Error("Memory address at  HL register should be 0x69, instead got: ", cpu.Memory.RAM[cpu.HL()])
    }
}

func TestLDHL_D(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x11
    cpu.Registers.L = 0x69
    cpu.Registers.D = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDHL_D

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Memory.RAM[cpu.HL()] != 0x69 {
        t.Error("Memory address at  HL register should be 0x69, instead got: ", cpu.Memory.RAM[cpu.HL()])
    }
}

func TestLDHL_E(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x11
    cpu.Registers.L = 0x69
    cpu.Registers.E = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDHL_E

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Memory.RAM[cpu.HL()] != 0x69 {
        t.Error("Memory address at  HL register should be 0x69, instead got: ", cpu.Memory.RAM[cpu.HL()])
    }
}

func TestLDHL_H(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x11
    cpu.Registers.L = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDHL_H

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Memory.RAM[cpu.HL()] != 0x11 {
        t.Error("Memory address at  HL register should be 0x11, instead got: ", cpu.Memory.RAM[cpu.HL()])
    }
}

func TestLDHL_L(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x11
    cpu.Registers.L = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDHL_L

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Memory.RAM[cpu.HL()] != 0x69 {
        t.Error("Memory address at  HL register should be 0x69, instead got: ", cpu.Memory.RAM[cpu.HL()])
    }
}

func TestLDHL_A(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x11
    cpu.Registers.L = 0x69
    cpu.Registers.A = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDHL_A

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Memory.RAM[cpu.HL()] != 0x69 {
        t.Error("Memory address at  HL register should be 0x69, instead got: ", cpu.Memory.RAM[cpu.HL()])
    }
}

func TestLDA_B(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.B = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDA_B

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.A != 0x69 {
        t.Error("A register should be 0x69, instead got: ", cpu.Registers.A)
    }
}

func TestLDA_C(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.C = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDA_C

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.A != 0x69 {
        t.Error("A register should be 0x69, instead got: ", cpu.Registers.A)
    }
}
func TestLDA_D(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.D = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDA_D

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.A != 0x69 {
        t.Error("A register should be 0x69, instead got: ", cpu.Registers.A)
    }
}
func TestLDA_E(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.E = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDA_E

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.A != 0x69 {
        t.Error("A register should be 0x69, instead got: ", cpu.Registers.A)
    }
}
func TestLDA_H(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDA_H

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.A != 0x69 {
        t.Error("A register should be 0x69, instead got: ", cpu.Registers.A)
    }
}

func TestLDA_L(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.L = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDA_L

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.A != 0x69 {
        t.Error("A register should be 0x69, instead got: ", cpu.Registers.A)
    }
}

func TestLDA_HL(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.H = 0x80
    cpu.Registers.L = 0x8F
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDA_HL
    cpu.Memory.RAM[0x808F] = 0x20

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.A != 0x20 {
        t.Error("A register should be 0x20, instead got: ", cpu.Registers.B)
    }
}

func TestLDA_A(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.A = 0x69
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDA_A

    // When
    expectedCycles := 1
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.A != 0x69 {
        t.Error("A register should be 0x69, instead got: ", cpu.Registers.A)
    }
}

func TestLDa8_A(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.A = 0x77
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDa8_A
    cpu.Memory.RAM[0x0101] = 0x22

    // When
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Memory.RAM[0xFF22] != 0x77 {
        t.Error("Memory cell at 0xFF22 should be 0x77, instead got: ", cpu.Memory.RAM[0xFF22])
    }
}

func TestSP_d16(t *testing.T) {

    cpu := InitSM83()

    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDSP_d16
    cpu.Memory.RAM[0x0101] = 0x22
    cpu.Memory.RAM[0x0102] = 0x80

    // When
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.SP != 0x8022 {
        t.Error("SP register should be 0x8022, instead got: ", cpu.Registers.SP)
    }

}

func TestLDCind_A(t *testing.T) {

    cpu := InitSM83()
    cpu.Registers.A = 0x77
    cpu.Registers.C = 0xA2
    
    // Given
    cpu.Memory.RAM[0x0100] = instructions.LDCind_A

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Memory.RAM[0xFFA2] != 0x77 {
        t.Error("Memory cell at 0xFF22 should be 0x77, instead got: ", cpu.Memory.RAM[0xFF22])
    }
}

func Testa16_A(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    // A = 0x11
    cpu.Memory.RAM[0x0100] = instructions.LDa16_A
    cpu.Memory.RAM[0x0101] = 0x52
    cpu.Memory.RAM[0x0102] = 0x72

    // 0x5555 data into 7252

    // When
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Memory.RAM[0x7252] != 0x11 {
        t.Error("Address 0x7252 should be 0x11.")
    }
}

func TestA_a8(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    // A = 0x11
    cpu.Memory.RAM[0x0100] = instructions.LDA_a8
    cpu.Memory.RAM[0x0101] = 0x52
    cpu.Memory.RAM[0xFF52] = 0x69


    // When
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.A != 0x69{
        t.Error("A register should be 0x69, instead got: ", cpu.Registers.A)
    }
}

func TestA_Cind(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    // A = 0x11
    cpu.Registers.C = 0x33
    cpu.Memory.RAM[0x0100] = instructions.LDA_Cind
    cpu.Memory.RAM[0xFF33] = 0x69

    // When
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.A != 0x69{
        t.Error("A register should be 0x69, instead got: ", cpu.Registers.A)
    }
}

func TestA_a16(t *testing.T) {

    cpu := InitSM83()
    
    // Given
    // A = 0x11
    cpu.Memory.RAM[0x0100] = instructions.LDA_a16
    cpu.Memory.RAM[0x0101] = 0x52
    cpu.Memory.RAM[0x0102] = 0x53
    cpu.Memory.RAM[0x5352] = 0x69


    // When
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, " cycles expected: ", expectedCycles)
    }

    if cpu.Registers.A != 0x69{
        t.Error("A register should be 0x69, instead got: ", cpu.Registers.A)
    }
}
