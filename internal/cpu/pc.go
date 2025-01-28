package cpu

type PC struct {
	// In
	counter uint32
	ImmExt  uint32
	AluRes  uint32
	// Control
	PCSrc int
}

func (pc *PC) compute() {
	switch pc.PCSrc {
	case 0:
		pc.counter += 4
	case 1:
		pc.counter += pc.ImmExt
	case 2:
		pc.counter = pc.AluRes
	}
}
