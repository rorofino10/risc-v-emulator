package cpu

const (
	INSTRUCTION_MEMORY_SIZE = 0xFFFFFFFF
)

type InstructionMemory struct {
	memory [INSTRUCTION_MEMORY_SIZE]byte
	A      uint32      // In
	RD     Instruction // Out
}

func (instr_mem *InstructionMemory) compute() {
	chunk1 := Instruction(instr_mem.memory[instr_mem.A])
	chunk2 := Instruction(instr_mem.memory[instr_mem.A+1])
	chunk3 := Instruction(instr_mem.memory[instr_mem.A+2])
	chunk4 := Instruction(instr_mem.memory[instr_mem.A+3])
	instr_mem.RD = chunk1 | (chunk2 << 8) | (chunk3 << 16) | (chunk4 << 24)
}

func (instr_mem *InstructionMemory) setInstructionAt(pos uint32, instr Instruction) {
	chunk1 := byte(instr & 0xFF)
	chunk2 := byte((instr >> 8) & 0xFF)
	chunk3 := byte((instr >> 16) & 0xFF)
	chunk4 := byte((instr >> 24) & 0xFF)
	instr_mem.memory[pos] = chunk1
	instr_mem.memory[pos+1] = chunk2
	instr_mem.memory[pos+2] = chunk3
	instr_mem.memory[pos+3] = chunk4
}

func (instr_mem *InstructionMemory) setInstructionsAt(pos uint32, instructions []Instruction) {
	for i, instr := range instructions {
		instr_mem.setInstructionAt(pos+uint32(i*4), instr)
	}
}

func (instr_mem *InstructionMemory) loadInstructions(instructions []Instruction) {
	instr_mem.setInstructionsAt(0, instructions)
}
