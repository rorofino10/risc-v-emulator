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

var c cpu.CPU = cpu.New()

func BenchmarkNOP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c.ExecuteInstruction(0x13)
	}
}
func BenchmarkADDNOP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c.ExecuteInstruction(0x33)
	}
}
func BenchmarkJUMPNOP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c.ExecuteInstruction(0x6f)
	}
}
