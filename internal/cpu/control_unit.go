package cpu

import (
	"errors"
)

const (
	lw_opcode    = 0b0000011
	sw_opcode    = 0b0100011
	r_opcode     = 0b0110011
	b_opcode     = 0b1100011
	i_opcode     = 0b0010011
	jal_opcode   = 0b1101111
	jal_r_opcode = 0b1100111

	I_ImmSrc = 0
	S_ImmSrc = 1
	B_ImmSrc = 2
	J_ImmSrc = 3
)

type ControlUnit struct {
	//In
	op       uint32
	funct3   uint32
	funct7_5 uint32
	Zero     bool
	//Out
	Branch      bool
	Jump        bool
	JumpFromALU bool
	ALUOp       int
	PCSrc       int
	ResultSrc   int
	MemWrite    bool
	ALUControl  int
	ALUSrc      int
	ImmSrc      int
	RegWrite    bool
}

func (control *ControlUnit) compute() error {
	switch control.op {
	case lw_opcode:
		control.RegWrite = true
		control.ImmSrc = I_ImmSrc
		control.ALUSrc = 1
		control.MemWrite = true
		control.ResultSrc = 1
		control.Branch = false
		control.ALUOp = 0b00
		control.Jump = false
	case sw_opcode:
		control.RegWrite = false
		control.ImmSrc = S_ImmSrc
		control.ALUSrc = 1
		control.MemWrite = true
		control.Branch = false
		control.ALUOp = 0b00
		control.Jump = false
	case r_opcode:
		control.RegWrite = true
		control.ALUSrc = 1
		control.MemWrite = false
		control.ResultSrc = 0
		control.Branch = false
		control.ALUOp = 0b10
		control.Jump = false
	case b_opcode:
		control.RegWrite = false
		control.ImmSrc = B_ImmSrc
		control.ALUSrc = 0
		control.MemWrite = false
		control.Branch = true
		control.ALUOp = 0b01
		control.Jump = false
	case i_opcode:
		control.RegWrite = true
		control.ImmSrc = I_ImmSrc
		control.ALUSrc = 1
		control.MemWrite = false
		control.ResultSrc = 0
		control.Branch = false
		control.Jump = false
		control.ALUOp = 0b10
	case jal_opcode:
		control.RegWrite = true
		control.ImmSrc = J_ImmSrc
		control.MemWrite = false
		control.ResultSrc = 2
		control.Branch = false
		control.Jump = true
		control.JumpFromALU = false
	case jal_r_opcode:
		control.RegWrite = true
		control.ImmSrc = I_ImmSrc
		control.ALUSrc = 1
		control.MemWrite = false
		control.ResultSrc = 2
		control.Branch = false
		control.Jump = true
		control.JumpFromALU = true
	default:
		return errors.New("invalid opcode")
	}

	switch control.ALUOp {
	case 0b00:
		control.ALUControl = 0b000 // lw, sw
	case 0b01:
		control.ALUControl = 0b001 // beq
	case 0b100: //  jal_r
		if control.funct3 == 0b000 {
			control.ALUControl = 0b000
		} else {
			return errors.New("invalid funct3")
		}
	case 0b10:
		switch control.funct3 {
		case 0b000:
			if control.op>>5 == 1 && control.funct7_5 == 1 {
				control.ALUControl = 0b001 // sub
			} else {
				control.ALUControl = 0b000 // add
			}
		case 0b010:
			control.ALUControl = 0b101 // slt
		case 0b110:
			control.ALUControl = 0b011 // or
		case 0b111:
			control.ALUControl = 0b010 // and
		default:
			return errors.New("invalid funct3")
		}
	}
	return nil
}

func (control *ControlUnit) computePCSrc() {
	if control.Jump && control.JumpFromALU {
		control.PCSrc = 2
	} else if control.Zero && control.Branch || control.Jump {
		control.PCSrc = 1
	} else {
		control.PCSrc = 0
	}
}
