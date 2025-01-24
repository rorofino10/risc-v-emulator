package cpu

type component interface {
	compute()
}

type memoryComponent interface {
	computeMemory()
}
