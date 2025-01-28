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
	for err := c.Execute(); err != nil; {

	}
}
