package cpu

import "errors"

type ALU struct {
	// In
	SrcA int32
	SrcB int32
	// Control
	AluControl uint8
	// Out
	AluResult int32
	Zero      bool
	Sign      bool // True is positive
}

func (alu *ALU) compute() error {
	switch alu.AluControl {
	case 0b000: // ADD
		alu.AluResult = alu.SrcA + alu.SrcB
	case 0b001: // SUB
		alu.AluResult = alu.SrcA - alu.SrcB
	case 0b010: // AND
		alu.AluResult = alu.SrcA & alu.SrcB
	case 0b011: // OR
		alu.AluResult = alu.SrcA | alu.SrcB
	case 0b101: // SLT
		alu.AluResult = (alu.SrcA - alu.SrcB) >> 31
	default:
		return errors.New("unknown alu control")
	}
	alu.Zero = alu.AluResult == 0
	alu.Sign = alu.AluResult>>31 == 0
	return nil
}
