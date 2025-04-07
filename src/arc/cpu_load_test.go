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


