package arc

import "testing"

func TestCPUResetsCorrectly(t *testing.T) {

    // Given
    cpu := InitSM83()

    want := &CPU{
        Registers: RegisterFile{
            PC: 0x0100,
            SP: 0xFFFE,
            A: 0x11,
            F: 0x80,
            B: 0x00,
            C: 0x00,
            D: 0xFF,
            E: 0x56,
            H: 0x00,
            L: 0x0D,
        } ,
    }
    if *cpu != *want {
        t.Error("CPU not reset correctly!")
    }
}

func TestCPUDoesNothingWhenExecutingZeroCycles(t *testing.T) {

    // Given
    const NUM_CYCLES = 0
    cpu := InitSM83()

    // When
    cyclesUsed := cpu.Execute(0)

    // Then
    if cyclesUsed != 0 {
        t.Error("Executing with zero cycles should return 0.")
    }
}


func InitSM83() (cpu *CPU){
    cpu = &CPU{}
    cpu.Memory = Memory{}
    cpu.ResetCPU()
    return
}
