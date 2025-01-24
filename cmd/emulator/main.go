package main

import (
	"github.com/rorofino10/risc-v-emulator/internal/cpu"
)

func main() {
	processor := cpu.New()
	var err error = nil
	for err == nil {
		err = processor.Execute()
	}
}
