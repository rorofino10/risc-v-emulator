package cpu

import "errors"

const (
	REGISTERS_AMOUNT = 32
)

type RegisterMemory struct {
	registers [REGISTERS_AMOUNT]uint32
	// Data Signal
	A1  uint8
	A2  uint8
	A3  uint8
	WD3 uint32
	// Control signal
	WE3 bool
	// Out
	RD1 uint32
	RD2 uint32
}

func (reg_mem *RegisterMemory) compute() {
	reg_mem.RD1 = reg_mem.registers[reg_mem.A1]
	reg_mem.RD2 = reg_mem.registers[reg_mem.A2]
}

func (reg_mem *RegisterMemory) computeMemory() {
	if reg_mem.WE3 {
		reg_mem.registers[reg_mem.A3] = reg_mem.WD3
	}
}

func (reg_mem *RegisterMemory) setRegister(reg int, data uint32) error {
	if reg >= len(reg_mem.registers) {
		return errors.New("invalid register")
	}
	reg_mem.registers[reg] = data
	return nil
}
