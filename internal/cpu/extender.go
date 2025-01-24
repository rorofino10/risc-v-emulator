package cpu

var sext_masks = [32]uint32{0xffffffff,
	0xfffffffe,
	0xfffffffc,
	0xfffffff8,
	0xfffffff0,
	0xffffffe0,
	0xffffffc0,
	0xffffff80,
	0xffffff00,
	0xfffffe00,
	0xfffffc00,
	0xfffff800,
	0xfffff000,
	0xffffe000,
	0xffffc000,
	0xffff8000,
	0xffff0000,
	0xfffe0000,
	0xfffc0000,
	0xfff80000,
	0xfff00000,
	0xffe00000,
	0xffc00000,
	0xff800000,
	0xff000000,
	0xfe000000,
	0xfc000000,
	0xf8000000,
	0xf0000000,
	0xe0000000,
	0xc0000000,
	0x80000000}

type Extender struct {
	// In
	Src uint32
	// Control
	ImmSrc int
	// Out
	ImmExt uint32
}

func sext(in uint32, n int) uint32 {
	return in | (sext_masks[n] * (in >> (n - 1)))
}

func (extender *Extender) compute() {
	switch extender.ImmSrc {
	case 0b00:
		newExt := extender.Src >> 20
		extender.ImmExt = sext(newExt, 12)
	case 0b01:
		newExt := (extender.Src >> 5) | (extender.Src >> 7 & 0x1f)
		extender.ImmExt = sext(newExt, 12)
	case 0b10:
		newExt := (extender.Src>>7&1)<<11 |
			(extender.Src>>31)<<12 |
			(((extender.Src >> 25) & 0x3F) << 5) |
			((extender.Src>>8)&0xF)<<1

		extender.ImmExt = sext(newExt, 13)
	}

}
