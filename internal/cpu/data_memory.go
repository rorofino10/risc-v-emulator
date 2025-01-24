package cpu

const (
	DATA_MEMORY_SIZE = 16384
)

type DataMemory struct {
	memory [DATA_MEMORY_SIZE]byte

	// In
	A  uint32
	WD uint32

	// Control flags
	WE bool

	// Out
	RD uint32
}

func (data_mem *DataMemory) compute() {
	chunk1 := data_mem.memory[data_mem.A]
	chunk2 := data_mem.memory[data_mem.A+1]
	chunk3 := data_mem.memory[data_mem.A+2]
	chunk4 := data_mem.memory[data_mem.A+3]

	data_mem.RD = uint32(chunk1) | (uint32(chunk2) << 8) | (uint32(chunk3) << 16) | (uint32(chunk4) << 24)
}

func (data_mem *DataMemory) storeWordAt(pos, word uint32) {
	chunk1 := byte(word & 0xFF)
	chunk2 := byte((word >> 8) & 0xFF)
	chunk3 := byte((word >> 16) & 0xFF)
	chunk4 := byte((word >> 24) & 0xFF)
	data_mem.memory[pos] = chunk1
	data_mem.memory[pos+1] = chunk2
	data_mem.memory[pos+2] = chunk3
	data_mem.memory[pos+3] = chunk4

}
