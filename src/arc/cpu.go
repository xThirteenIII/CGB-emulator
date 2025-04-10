package arc

import (
	"cgbemu/src/instructions"
	"fmt"
	"log"
)

// Following these two docs for Instruction implementation:
//
// https://gekkio.fi/files/gb-docs/gbctr.pdf, for detailed operations.
// https://meganesu.github.io/generate-gb-opcodes/

// CGB memory goes from 0x0000 to 0xFFFF.
const MaxMem = 1024 * 64

// 8-bit data bus, 16-bit address bus (output only).
type Memory struct {
    RAM    [MaxMem]byte
}

// Init initializes the memory to zero.
func (m *Memory) ClearRAM(){
    for i:=0; i<MaxMem; i++ {
        m.RAM[i] = 0
    }
}

// ResetCPU clears RAM (everything to 0) and loads initial values to registers.
func (cpu *CPU) ResetCPU() {
    cpu.Memory.ClearRAM()
    cpu.Registers.InitRegisters()
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

    F byte      // Flags

    // Special purpose 8-bit registers.
    IR byte     // Interrupt Register.
    IE byte     // Interrupt Enable.
}

// Initial registers values depend on the GameBoy Model.
// For the CGB mode:
// AF -> 0x1180
// BC -> 0x0000
// DE -> 0xFF56
// HL -> 0x000D
// SP -> 0xFFFE
// PC -> 0x0100
// We are skipping the Boot ROM istructions.
func (r *RegisterFile) InitRegisters() {

    r.A = byte(0b00010001) // 0x11
    r.F = byte(0b10000000) // 0x80

    r.B = byte(0x00)
    r.C = byte(0x00)

    r.D = byte(0xFF)
    r.E = byte(0x56)

    r.H = byte(0x00)
    r.L = byte(0x0D)

    r.SP = 0xFFFE

    r.PC = 0x0100
}

// FlagRegister contains informations about the last instruction that affected the flags.
// The Zero Flag and the Carry Flag are used for conditional instructions.
// The Carry flag is also used by arithmetic and logic instructions. The BCD Flags are used only by DAA instruction.
//
// The F register can't be accessed normally, only by doing a “push af/pop bc”, for example. The lower
// four bits are always zero, even if a “pop af” instruction tries to write other values.
/*
    Z uint  // Bit 7: Zero Flag.
    N uint  // Bit 6: Add/Sub Flag (BCD).
    H uint  // Bit 5: Half Carry Flag (BCD).
    C uint  // Bit 4: Carry Flag.
    U3 uint // Unused (always zero).
    U2 uint // Unused (always zero).
    U1 uint // Unused (always zero).
    U0 uint // Unused (always zero).
*/
// To access the F register, use bitwise operations.
// GET Flag:    & 0x80, &0x40, &0x20, &0x10 for Bit 7, Bit 6, Bit 5 , Bit 4.
// SET Flag:    |= 0x80, |= 0x40, |= 0x20, |= 0x10. 1 | 1 = 1, 0 | 1 = 1 so it does set nonetheless.
// CLEAR Flag:  &= ^0x80 and so on...
// ^ operator INVERTS all the bits.
type CPU struct {

    Memory      Memory

    Registers   RegisterFile

    IDU uint16
}

// PrintStatus prints registers values on Stdout.
func (cpu *CPU) PrintStatus() {
    fmt.Println("PC: ",     cpu.Registers.PC)
    fmt.Println("SP: ",     cpu.Registers.SP)
    fmt.Println("Flags: ",  cpu.Registers.F)
}

// Execute runs the fetch-decode loop.
// It fetches the instruction byte and then, based on the opcode fetched, 
// executes the corresponding instruction.
// It returns the number of cycles used, for Testing purposes.
func (cpu *CPU) Execute(cycles int) (cyclesUsed int) {

    cyclesUsed = cycles
    
    // Can we get stuck in infinite loop if we pass more cycles than expected?
    // Not for now because since memory is initialised to 0, if we try to fetch a 
    // byte from one more cell memory where we are not supposed to be, it fetches 0 and
    // exits the switch loop with the default case.
    for cycles > 0 {

        // For each byte of the current instrunction length, a FetchByte() operation is needed.
        //
        // Read opcode, 1 cycle used.
        ins := cpu.FetchByte(&cycles)

        // Decode instruction.
        switch ins {
            
        case instructions.LDB_IM: // Load to the 8-bit register B, the immediate data n.

            // FetchByte takes up one cycle.
            n8 := cpu.FetchByte(&cycles)
            cpu.Registers.B = n8

            // Length: 2 bytes, opcode + n.
            // Cycles: 2 machine cycles.
        case instructions.LDB_HL: // Load to the 8-bit register B, data from the absolute address specified by the 16-bit register HL.

            cpu.Registers.B = cpu.ReadByteFromMemory(&cycles, cpu.HL())

            // Length: 1 byte, opcode.
            // Cycles: 2 machine cycles.
        case instructions.LDBC_d16: // Load the 2 bytes of immediate data into register pair BC. Little-endian.

            // Read LSB from memory.
            nn_lsb := cpu.FetchByte(&cycles)

            // Read MSB from memory.
            nn_msb := cpu.FetchByte(&cycles)

            cpu.Registers.B = nn_msb
            cpu.Registers.C = nn_lsb
            // Length: 3 bytes, opcode + LSB(nn) + MSB(nn).
            // Cycles: 3 machine cycles.
        case instructions.LDBC_A: // Load to the absolute address specified by the 16-bit register BC, data from the 8-bit A register.

            cpu.WriteByteToMemory(&cycles, cpu.BC(), cpu.Registers.A)

            // Length: 1 byte, opcode.
            // Cycles: 2 machine cycles.
        case instructions.LDa16_SP: // Load to the absolute address specified by the 16-bit operand nn, data from the 16-bit SP register.

            // Read address lsb
            nn_lsb := cpu.FetchByte(&cycles)
            // Read address msb
            nn_msb := cpu.FetchByte(&cycles)

            // Compose absolute address.
            nn := GetUint16AddressFromLSBAndMSB(nn_lsb, nn_msb)

            // Write Stack Pointer LSB first.
            cpu.WriteByteToMemory(&cycles, nn, byte(cpu.Registers.SP & 0xFF))

            // Increment address by 1.
            nn = nn+1

            // Write Stack Pointer MSB last.
            cpu.WriteByteToMemory(&cycles, nn, byte(cpu.Registers.SP >> 8))

            // Length: 3 bytes, opcode + lsb + msb.
            // Cycles: 5 machine cycles. opcode, R, R, W, W.
        case instructions.LDA_BC:// Load to the 8-bit A register, data from the absolute address specified by the 16-bit register BC.

            cpu.Registers.A = cpu.ReadByteFromMemory(&cycles, cpu.BC())
            // Length: 1 byte.
            // Cycles: 2 machine cycles. opcode, R
        case instructions.LDC_d8: // Load the 8-bit immediate operand d8 into register C.
            
            cpu.Registers.C = cpu.FetchByte(&cycles)

            // Length: 2 bytes, opcode + n.
            // Cycles: 2 machine cycles. opcode, R
        case instructions.LDDE_d16: // Load the 2 bytes of immediate data into register pair DE. Little-endian.

            // Read LSB from memory.
            nn_lsb := cpu.FetchByte(&cycles)

            // Read MSB from memory.
            nn_msb := cpu.FetchByte(&cycles)

            cpu.Registers.D = nn_msb
            cpu.Registers.E = nn_lsb
            // Length: 3 bytes, opcode + LSB(nn) + MSB(nn).
            // Cycles: 3 machine cycles.
        case instructions.LDDE_A:   // Load to the absolute address specified by the 16-bit register DE, data from the 8-bit A register.

            cpu.WriteByteToMemory(&cycles, cpu.DE(), cpu.Registers.A)

            // Length: 1 byte.
            // Cycles: 2 machine cycles. opcode, R.
        case instructions.LDD_d8:   // Load to the absolute address specified by the 16-bit register DE, data from the 8-bit A register.

            cpu.Registers.D = cpu.FetchByte(&cycles)
            // Length: 2 bytes, opcode + n.
            // Cycles: 2 machine cycles. opcode, R.
        case instructions.LDA_DE:// Load to the 8-bit A register, data from the absolute address specified by the 16-bit register DE.

            cpu.Registers.A = cpu.ReadByteFromMemory(&cycles, cpu.DE())
            // Length: 1 byte.
            // Cycles: 2 machine cycles. opcode, R
        case instructions.LDE_d8: // Load the 8-bit immediate operand d8 into register E.
            
            cpu.Registers.E = cpu.FetchByte(&cycles)
            // Length: 2 bytes, opcode + n.
            // Cycles: 2 machine cycles. opcode, R
        case instructions.LDHL_d16: // Load the 2 bytes of immediate data into register pair HL. Little-endian.

            // Read LSB from memory.
            nn_lsb := cpu.FetchByte(&cycles)

            // Read MSB from memory.
            nn_msb := cpu.FetchByte(&cycles)

            cpu.Registers.H = nn_msb
            cpu.Registers.L = nn_lsb
            // Length: 3 bytes, opcode + LSB(nn) + MSB(nn).
            // Cycles: 3 machine cycles.
        case instructions.LDHLinc_A: // Load to the absolute address specified by the 16-bit register HL, data from the 8-bit A register.
                                    // The value of HL is incremented after the memory write.

            cpu.WriteByteToMemory(&cycles, cpu.HL(), cpu.Registers.A)
            Increment16Address(&cpu.Registers.L, &cpu.Registers.H)
            // Length: 1 byte.
            // Cycles: 2 machine cycles. opcode + W.
        case instructions.LDH_d8:   // Load the 8-bit immediate operand d8 into register H.

            cpu.Registers.H = cpu.FetchByte(&cycles)
            // Length: 2 bytes, opcode + n.
            // Cycles: 2 machine cycles. opcode, R.
        case instructions.LDA_HLinc:   // Load to the 8-bit A register, data from the absolute address specified by the 16-bit register HL.
                                      // The value of HL is incremented after the memory read.

            cpu.Registers.A = cpu.ReadByteFromMemory(&cycles, cpu.HL())
            Increment16Address(&cpu.Registers.L, &cpu.Registers.H)
            // Length: 1 byte.
            // Cycles: 2 machine cycles. opcode + W.
        case instructions.LDL_d8:   // Load the 8-bit immediate operand d8 into register L.

            cpu.Registers.L = cpu.FetchByte(&cycles)
            // Length: 2 bytes, opcode + n.
            // Cycles: 2 machine cycles. opcode, R.
        case instructions.LDHLdec_A: // Load to the absolute address specified by the 16-bit register HL, data from the 8-bit A register.
                                    // The value of HL is decremented after the memory write.

            cpu.WriteByteToMemory(&cycles, cpu.HL(), cpu.Registers.A)
            Decrement16Address(&cpu.Registers.L, &cpu.Registers.H)
            // Length: 1 byte.
            // Cycles: 2 machine cycles. opcode + W.
        case instructions.LDHL_d8:   // Load to the absolute address specified by the 16-bit register HL, the immediate data n.
            operand := cpu.FetchByte(&cycles)
            cpu.WriteByteToMemory(&cycles, cpu.HL(), operand)
            // Length: 2 bytes, opcode + n.
            // Cycles: 3 machine cycles. opcode, R, W.
        case instructions.LDA_HLdec:   // Load to the 8-bit A register, data from the absolute address specified by the 16-bit register HL.
                                      // The value of HL is decremented after the memory read.

            cpu.Registers.A = cpu.ReadByteFromMemory(&cycles, cpu.HL())
            Decrement16Address(&cpu.Registers.L, &cpu.Registers.H)
            // Length: 1 byte.
            // Cycles: 2 machine cycles. opcode + W.
        case instructions.LDA_d8:   // Load the 8-bit immediate operand d8 into register A.

            cpu.Registers.A = cpu.FetchByte(&cycles)
            // Length: 2 bytes, opcode + n.
            // Cycles: 2 machine cycles. opcode, R.
        case instructions.LDB_B:
            cpu.Registers.B = cpu.Registers.B
        case instructions.LDB_C:
            cpu.Registers.B = cpu.Registers.C
        case instructions.LDB_D:
            cpu.Registers.B = cpu.Registers.D
        case instructions.LDB_E:
            cpu.Registers.B = cpu.Registers.E
        case instructions.LDB_H:
            cpu.Registers.B = cpu.Registers.H
        case instructions.LDB_L:
            cpu.Registers.B = cpu.Registers.L
        case instructions.LDB_A:
            cpu.Registers.B = cpu.Registers.A   // Lenght: 1 byte.
                                                // Cycles: 1 cycle.
        case instructions.LDC_B:
            cpu.Registers.C = cpu.Registers.B
        case instructions.LDC_C:
            cpu.Registers.C = cpu.Registers.C
        case instructions.LDC_D:
            cpu.Registers.C = cpu.Registers.D
        case instructions.LDC_E:
            cpu.Registers.C = cpu.Registers.E
        case instructions.LDC_H:
            cpu.Registers.C = cpu.Registers.H
        case instructions.LDC_L:
            cpu.Registers.C = cpu.Registers.L
        case instructions.LDC_HL:
            cpu.Registers.C = cpu.ReadByteFromMemory(&cycles, cpu.HL())
        case instructions.LDC_A:
            cpu.Registers.C = cpu.Registers.A   // Lenght: 1 byte.
                                                // Cycles: 1 cycle.
        case instructions.LDD_B:
            cpu.Registers.D = cpu.Registers.B
        case instructions.LDD_C:
            cpu.Registers.D = cpu.Registers.C
        case instructions.LDD_D:
            cpu.Registers.D = cpu.Registers.D
        case instructions.LDD_E:
            cpu.Registers.D = cpu.Registers.E
        case instructions.LDD_H:
            cpu.Registers.D = cpu.Registers.H
        case instructions.LDD_L:
            cpu.Registers.D = cpu.Registers.L
        case instructions.LDD_HL:
            cpu.Registers.D = cpu.ReadByteFromMemory(&cycles, cpu.HL())
        case instructions.LDD_A:
            cpu.Registers.D = cpu.Registers.A   // Lenght: 1 byte.
                                                // Cycles: 1 cycle.
        case instructions.LDE_B:
            cpu.Registers.E = cpu.Registers.B
        case instructions.LDE_C:
            cpu.Registers.E = cpu.Registers.C
        case instructions.LDE_D:
            cpu.Registers.E = cpu.Registers.D
        case instructions.LDE_E:
            cpu.Registers.E = cpu.Registers.E
        case instructions.LDE_H:
            cpu.Registers.E = cpu.Registers.H
        case instructions.LDE_L:
            cpu.Registers.E = cpu.Registers.L
        case instructions.LDE_HL:
            cpu.Registers.E = cpu.ReadByteFromMemory(&cycles, cpu.HL())
        case instructions.LDE_A:
            cpu.Registers.E = cpu.Registers.A   // Lenght: 1 byte.
                                                // Cycles: 1 cycle.
        case instructions.LDH_B:
            cpu.Registers.H = cpu.Registers.B
        case instructions.LDH_C:
            cpu.Registers.H = cpu.Registers.C
        case instructions.LDH_D:
            cpu.Registers.H = cpu.Registers.D
        case instructions.LDH_E:
            cpu.Registers.H = cpu.Registers.E
        case instructions.LDH_H:
            cpu.Registers.H = cpu.Registers.H
        case instructions.LDH_L:
            cpu.Registers.H = cpu.Registers.L
        case instructions.LDH_HL:
            cpu.Registers.H = cpu.ReadByteFromMemory(&cycles, cpu.HL())
        case instructions.LDH_A:
            cpu.Registers.H = cpu.Registers.A   // Lenght: 1 byte.
                                                // Cycles: 1 cycle.
        case instructions.LDL_B:
            cpu.Registers.L = cpu.Registers.B
        case instructions.LDL_C:
            cpu.Registers.L = cpu.Registers.C
        case instructions.LDL_D:
            cpu.Registers.L = cpu.Registers.D
        case instructions.LDL_E:
            cpu.Registers.L = cpu.Registers.E
        case instructions.LDL_H:
            cpu.Registers.L = cpu.Registers.H
        case instructions.LDL_L:
            cpu.Registers.L = cpu.Registers.L
        case instructions.LDL_HL:
            cpu.Registers.L = cpu.ReadByteFromMemory(&cycles, cpu.HL())
        case instructions.LDL_A:
            cpu.Registers.L = cpu.Registers.A   // Lenght: 1 byte.
                                                // Cycles: 1 cycle.
        case instructions.LDHL_B:
            cpu.WriteByteToMemory(&cycles, cpu.HL(), cpu.Registers.B)
        case instructions.LDHL_C:
            cpu.WriteByteToMemory(&cycles, cpu.HL(), cpu.Registers.C)
        case instructions.LDHL_D:
            cpu.WriteByteToMemory(&cycles, cpu.HL(), cpu.Registers.D)
        case instructions.LDHL_E:
            cpu.WriteByteToMemory(&cycles, cpu.HL(), cpu.Registers.E)
        case instructions.LDHL_H:
            cpu.WriteByteToMemory(&cycles, cpu.HL(), cpu.Registers.H)
        case instructions.LDHL_L:
            cpu.WriteByteToMemory(&cycles, cpu.HL(), cpu.Registers.L)
        case instructions.LDHL_A:
            cpu.WriteByteToMemory(&cycles, cpu.HL(), cpu.Registers.A)
                                                // Lenght: 2 bytes.
                                                // Cycles: 2 cycle.
        case instructions.LDA_B:
            cpu.Registers.A = cpu.Registers.B
        case instructions.LDA_C:
            cpu.Registers.A = cpu.Registers.C
        case instructions.LDA_D:
            cpu.Registers.A = cpu.Registers.D
        case instructions.LDA_E:
            cpu.Registers.A = cpu.Registers.E
        case instructions.LDA_H:
            cpu.Registers.A = cpu.Registers.H
        case instructions.LDA_L:
            cpu.Registers.A = cpu.Registers.L
        case instructions.LDA_HL:
            cpu.Registers.A = cpu.ReadByteFromMemory(&cycles, cpu.HL())
        case instructions.LDA_A:
            cpu.Registers.A = cpu.Registers.A   // Lenght: 1 byte.
                                                // Cycles: 1 cycle.
        case instructions.LDa8_A:
            //Load to the address specified by the 8-bit immediate data n, data from the 8-bit A register. The
            //full 16-bit absolute address is obtained by setting the most significant byte to 0xFF and the
            //least significant byte to the value of n, so the possible range is 0xFF00-0xFFFF.
            //
            //Should specify a 16-bit address in the mnemonic portion for a8, although the immediate operand only has the lower-order 8 bits.

            //0xFF00-0xFF7F: Port/Mode registers, control register, sound register
            //0xFF80-0xFFFE: Working & Stack RAM (127 bytes)
            //0xFFFF: Interrupt Enable Register

            // Length: 2 bytes, opcode+n.
            // Cycles: 3 cycles, opcode + R + W
            n := cpu.FetchByte(&cycles)
            absoluteAddress := 0xFF00 | uint16(n)
            cpu.WriteByteToMemory(&cycles, absoluteAddress, cpu.Registers.A)
        case instructions.LDSP_d16: // Load to the 16-bit register rr, the immediate 16-bit data nn.

            nn_lsb := cpu.FetchByte(&cycles)
            nn_msb := cpu.FetchByte(&cycles)

            nn := GetUint16AddressFromLSBAndMSB(nn_lsb, nn_msb)

            cpu.Registers.SP = nn

            // Length: 3 bytes, opcode+lsb(n)+msb(n).
            // Cycles: 3 cycles, opcode + R + R
        case instructions.LDCind_A: //Load to the address specified by the 8-bit C register, data from the 8-bit A register. The full
                                    //16-bit absolute address is obtained by setting the most significant byte to 0xFF and the least
                                    //significant byte to the value of C, so the possible range is 0xFF00-0xFFFF.

            absoluteAddress := 0xFF00 | uint16(cpu.Registers.C)
            cpu.WriteByteToMemory(&cycles, absoluteAddress, cpu.Registers.A)

            // Length: 1 bytes, opcode
            // Cycles: 2 cycles, opcode + R
        case instructions.LDa16_A: // Load to the absolute address specified by the 16-bit operand nn, data from the A register.

            // Read address lsb
            nn_lsb := cpu.FetchByte(&cycles)
            // Read address msb
            nn_msb := cpu.FetchByte(&cycles)

            // Compose absolute address.
            nn := GetUint16AddressFromLSBAndMSB(nn_lsb, nn_msb)

            // Write Stack Pointer MSB last.
            cpu.WriteByteToMemory(&cycles, nn, cpu.Registers.A)

            // Length: 3 bytes, opcode + lsb + msb.
            // Cycles: 4 machine cycles. opcode, R, R, W.
        case instructions.LDA_a8:
            // Load to the 8-bit A register, data from the address specified by the 8-bit immediate data n. The
            // full 16-bit absolute address is obtained by setting the most significant byte to 0xFF and the
            // least significant byte to the value of n, so the possible range is 0xFF00-0xFFFF.

            n := cpu.FetchByte(&cycles)

            absoluteAddress := 0xFF00 | uint16(n)
            cpu.Registers.A = cpu.ReadByteFromMemory(&cycles, absoluteAddress)
            
            // Length: 2 bytes, opcode + n
            // Cycles: 3 machine cycles. opcode, R, R.
        case instructions.LDA_Cind:
            // Load to the 8-bit A register, data from the address specified by the 8-bit C register. The full
            // 16-bit absolute address is obtained by setting the most significant byte to 0xFF and the least
            // significant byte to the value of C, so the possible range is 0xFF00-0xFFFF.
            absoluteAddress := 0xFF00 | uint16(cpu.Registers.C)
            cpu.Registers.A = cpu.ReadByteFromMemory(&cycles, absoluteAddress)

            // Length: 1 bytes, opcode.
            // Cycles: 2 machine cycles. opcode, R.
        case instructions.LDA_a16:
            // Load to the 8-bit A register, data from the address specified by the 16-bit immediate data n.

            nn_lsb := cpu.FetchByte(&cycles)
            nn_msb := cpu.FetchByte(&cycles)

            absoluteAddress := GetUint16AddressFromLSBAndMSB(nn_lsb, nn_msb)
            cpu.Registers.A = cpu.ReadByteFromMemory(&cycles, absoluteAddress)
            // Length: 3 bytes, opcode + lsb + msb
            // Cycles: 4 machine cycles. opcode, R(lsb), R(msb), R(absAddr).
        case instructions.POP_BC:
            // Pops to the 16-bit register rr, data from the stack memory.
            // This instruction does not do calculations that affect flags, but POP AF completely replaces the
            // F register value, so all flags are changed based on the 8-bit data that is read from memory.
            lsb := cpu.PopFromSP(&cycles)
            msb := cpu.PopFromSP(&cycles)

            cpu.Registers.B = msb
            cpu.Registers.C = lsb
            // Length: 1 byte
            // Cycles: 3 machine cycles. opcode, R(lsb), R(msb)
        case instructions.PUSH_BC:
            // Push to the stack memory, data from the 16-bit register BC.
            //
            // Push MSB first, id est B register.
            // Since SP grows downward, msb is read first?
            cpu.Registers.SP--
            cycles-- // A cycle is consumed just for decrementing SP.
            cpu.WriteByteToMemory(&cycles, cpu.Registers.SP, cpu.Registers.B)
            cpu.Registers.SP--
            cpu.WriteByteToMemory(&cycles, cpu.Registers.SP, cpu.Registers.C)

            // Length: 1 byte
            // Cycles: 4 machine cycles. opcode, W(lsb), W(msb), 
        case instructions.POP_DE:
            // Pops to the 16-bit register rr, data from the stack memory.
            // This instruction does not do calculations that affect flags, but POP AF completely replaces the
            // F register value, so all flags are changed based on the 8-bit data that is read from memory.
            lsb := cpu.PopFromSP(&cycles)
            msb := cpu.PopFromSP(&cycles)

            cpu.Registers.D = msb
            cpu.Registers.E = lsb
            // Length: 1 byte
            // Cycles: 3 machine cycles. opcode, R(lsb), R(msb)
        case instructions.PUSH_DE:
            // Push to the stack memory, data from the 16-bit register DE.
            //
            // Push MSB first, id est D register.
            // Since SP grows downward, msb is read first?
            cpu.Registers.SP--
            cycles-- // A cycle is consumed just for decrementing SP.
            cpu.WriteByteToMemory(&cycles, cpu.Registers.SP, cpu.Registers.D)
            cpu.Registers.SP--
            cpu.WriteByteToMemory(&cycles, cpu.Registers.SP, cpu.Registers.E)

            // Length: 1 byte
            // Cycles: 4 machine cycles. opcode, W(lsb), W(msb), 
        case instructions.POP_HL:
            // Pops to the 16-bit register rr, data from the stack memory.
            // This instruction does not do calculations that affect flags, but POP AF completely replaces the
            // F register value, so all flags are changed based on the 8-bit data that is read from memory.
            lsb := cpu.PopFromSP(&cycles)
            msb := cpu.PopFromSP(&cycles)

            cpu.Registers.H = msb
            cpu.Registers.L = lsb
            // Length: 1 byte
            // Cycles: 3 machine cycles. opcode, R(lsb), R(msb)
        case instructions.PUSH_HL:
            // Push to the stack memory, data from the 16-bit register HL.
            //
            // Push MSB first, id est H register.
            // Since SP grows downward, msb is read first?
            cpu.Registers.SP--
            cycles-- // A cycle is consumed just for decrementing SP.
            cpu.WriteByteToMemory(&cycles, cpu.Registers.SP, cpu.Registers.H)
            cpu.Registers.SP--
            cpu.WriteByteToMemory(&cycles, cpu.Registers.SP, cpu.Registers.L)

            // Length: 1 byte
            // Cycles: 4 machine cycles. opcode, W(lsb), W(msb), 
        case instructions.POP_AF:
            // Pops to the 16-bit register rr, data from the stack memory.
            // POP AF completely replaces the F register value, so all flags are changed based on the 8-bit data that is read from memory.
            lsb := cpu.PopFromSP(&cycles)
            msb := cpu.PopFromSP(&cycles)

            cpu.Registers.A = msb
            cpu.Registers.F = lsb
            // Length: 1 byte
            // Cycles: 3 machine cycles. opcode, R(lsb), R(msb)
        case instructions.PUSH_AF:
            // Push to the stack memory, data from the 16-bit register AF.
            //
            // Push MSB first, id est B register.
            // Since SP grows downward, msb is read first?
            cpu.Registers.SP--
            cycles-- // A cycle is consumed just for decrementing SP.
            cpu.WriteByteToMemory(&cycles, cpu.Registers.SP, cpu.Registers.A)
            cpu.Registers.SP--
            cpu.WriteByteToMemory(&cycles, cpu.Registers.SP, cpu.Registers.F)

            // Length: 1 byte
            // Cycles: 4 machine cycles. opcode, W(lsb), W(msb), 
        default:

            log.Println("At memory address: ", cpu.Registers.PC)

            // TODO: Should it stop and Fatal or just keep going till next valid instruction?
            log.Fatalln("Unknown opcode: ", ins)}
    }

    // If the number of cycles used is correct, respectively to the instruction used, 
    // the return should be the original value, passed when calling Execute().
    // When testing the instruction, we make sure that the expected value returned by Execute()
    // matches the cycles needed for the instructions, based on official documentation.
    cyclesUsed -= cycles
    
    // e.g. cpu.Execute(2) when executing 0x06 (LDB_IM):
    //      cyclesUsed = cycles = 2
    //      executing LDB_IM consumes 2 machine cycles.
    //      cycles becomes 0.
    //      cyclesUsed = cyclesUsed - 0
    //      return
    return
}
