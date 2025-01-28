package cpu

import "fmt"

type Instruction uint32

type CPU struct {
	pc           PC
	Clock        chan struct{}
	instr_mem    InstructionMemory
	control_unit ControlUnit
	extender     Extender
	alu          ALU
	reg_mem      RegisterMemory
	data_mem     DataMemory
}

func New() *CPU {
	cpu := &CPU{
		pc:           PC{},
		Clock:        make(chan struct{}),
		instr_mem:    InstructionMemory{},
		control_unit: ControlUnit{},
		extender:     Extender{},
		alu:          ALU{},
		reg_mem:      RegisterMemory{},
		data_mem:     DataMemory{},
	}
	// instructions := []instruction{0x00128293, 0xffdff0ef}

	return cpu
}

func (c *CPU) LoadInstructions(instructions []Instruction) {
	c.instr_mem.loadInstructions(instructions)
}

func (c *CPU) Run() {
	for range c.Clock {
		fmt.Print("\033[H\033[2J")
		for i, v := range c.reg_mem.registers {
			fmt.Printf("[x%d]:%x\n", i, v)
		}
		c.Execute()
	}
}

func (c *CPU) Execute() error {
	// Get Instruction using PC from Instruction Memory
	c.instr_mem.A = c.pc.counter
	c.instr_mem.compute()
	instr := c.instr_mem.RD

	// Control Flags
	c.control_unit.op = uint32(instr) & 0x7F
	c.control_unit.funct3 = uint32(instr) >> 12 & 0x7
	c.control_unit.funct7_5 = uint32(instr) >> 30 & 1
	c.control_unit.compute()

	// Handle Register Component
	c.reg_mem.A1 = uint8(instr >> 15 & 0x1F)
	c.reg_mem.A2 = uint8(instr >> 20 & 0x1F)
	c.reg_mem.A3 = uint8(instr >> 7 & 0x1F)
	c.reg_mem.WE3 = c.control_unit.RegWrite
	c.reg_mem.compute()

	c.extender.Src = uint32(instr)
	c.extender.ImmSrc = c.control_unit.ImmSrc
	c.extender.compute()

	c.alu.SrcA = int32(c.reg_mem.RD1)
	switch c.control_unit.ALUSrc {
	case 0:
		c.alu.SrcB = int32(c.reg_mem.RD2)
	case 1:
		c.alu.SrcB = int32(c.extender.ImmExt)

	}
	c.alu.AluControl = uint8(c.control_unit.ALUControl)
	c.alu.compute()
	c.control_unit.Zero = c.alu.Zero

	c.data_mem.A = uint32(c.alu.AluResult)
	fmt.Println(c.alu.AluResult)
	fmt.Println(c.data_mem.A)
	c.data_mem.WD = c.reg_mem.RD2
	c.data_mem.WE = c.control_unit.MemWrite
	c.data_mem.compute()

	// Write data to Register
	switch c.control_unit.ResultSrc {
	case 0:
		c.reg_mem.WD3 = uint32(c.alu.AluResult)
	case 1:
		c.reg_mem.WD3 = c.data_mem.RD
	case 2:
		c.reg_mem.WD3 = c.pc.counter + 4
	}
	c.reg_mem.computeMemory()

	c.control_unit.computePCSrc()
	c.pc.ImmExt = c.extender.ImmExt
	c.pc.AluRes = uint32(c.alu.AluResult)
	c.pc.PCSrc = c.control_unit.PCSrc
	c.pc.compute()
	return nil
}
