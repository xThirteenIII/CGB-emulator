package arc

// CGB memory goes from 0x0000 to 0xFFFF.
const MaxMem = 1024 * 64

type Memory struct {
    Data    [MaxMem]byte
}

// Init initializes the memory to zero.
func (m *Memory) Init(){
    for i:=0; i<MaxMem; i++ {
        m.Data[i] = 0
    }
}

// The CGB CPU is an 8-bit 8080-like Sharp CPU (speculated to be a SM83 core).
//
// Main sub-systems of a SM83:
// * Control Unit
// * Register file
// * ALU
// * IDU
type RegisterFile struct {

    // Program Counter points to the next instruction.
    PC uint16

    // Stack Pointer holds the location of the next free bytes on the stack.
    SP uint16

    // The CGB has 8 bit and 16 bit registers,
    // and some of the 8 bit registers can be used together to form a 16-bit value and use it like a number.
    A byte      // Accumulator & Flags.
    B byte      // 
    C byte      // BC
    D byte      // 
    E byte      // DE
    H byte      //
    L byte      // HL

    F FlagRegister      // Flags

    // Special purpose 8-bit registers.
    IR byte     // Interrupt Register.
    IE byte     // Interrupt Enable.
}

// FlagRegister contains informations about the last instruction that affected the flags.
// The Zero Flag and the Carry Flag are used for conditional instructions.
// The Carry flag is also used by arithmetic and logic instructions. The BCD Flags are used only by DAA instruction.
//
// The F register can't be accessed normally, only by doing a “push af/pop bc”, for example. The lower
// four bits are always zero, even if a “pop af” instruction tries to write other values.
type FlagRegister struct {

    Z uint  // Bit 7: Zero Flag.
    N uint  // Bit 6: Add/Sub Flag (BCD).
    H uint  // Bit 5: Half Carry Flag (BCD).
    C uint  // Bit 4: Carry Flag.
    U3 uint // Unused (always zero).
    U2 uint // Unused (always zero).
    U1 uint // Unused (always zero).
    U0 uint // Unused (always zero).
}

type CPU struct {

    Memory      Memory

    Registers   RegisterFile

    IDU uint16
}
