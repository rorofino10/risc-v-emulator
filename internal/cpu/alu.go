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
}

func (alu *ALU) compute() error {
	switch alu.AluControl {
	case 0b000:
		alu.AluResult = alu.SrcA + alu.SrcB
	case 0b001:
		alu.AluResult = alu.SrcA - alu.SrcB
	case 0b0101:
		return errors.New("missing slt")
	case 0b011:
		alu.AluResult = alu.SrcA | alu.SrcB
	case 0b010:
		alu.AluResult = alu.SrcA & alu.SrcB
	default:
		return errors.New("unknown alu op code")
	}
	alu.Zero = alu.AluResult == 0
	return nil
}
