package cpu_test

import (
	"testing"

	"github.com/rorofino10/risc-v-emulator/internal/cpu"
)

func TestSampleProgram(t *testing.T) {
	c := cpu.New()
	instructions := []cpu.Instruction{
		0xffc00513,
		0x000525b3,
		0x004000e7,
	}
	c.LoadInstructions(instructions)
	var err error = nil
	for err == nil {
		err = c.Execute()
	}
}

var c *cpu.CPU = cpu.New().LoadInstructions([]cpu.Instruction{
	0x00000013,
})

func BenchmarkAddi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c.Execute()
	}
}
