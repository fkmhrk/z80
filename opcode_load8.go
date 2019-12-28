package z80

var load8 = []OPCode{

	OPCode{
		N: "LD r1, r2",
		C: []Code{
			{0x40, 0x3f},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			*cpu.regP(codes[0] >> 3) = *cpu.regP(codes[0])
		},
	},

	OPCode{
		N: "LD r, n",
		C: []Code{
			{0x05, 0x38},
			{0x00, 0xff},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			r := cpu.regP(codes[0] >> 3)
			*r = codes[1]
		},
	},

	OPCode{
		N: "LD r, (HL)",
		C: []Code{
			{0x45, 0x38},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			r := cpu.regP(codes[0] >> 3)
			*r = cpu.Memory.Get(cpu.HL.U16())
		},
	},

	OPCode{
		N: "LD r, (IX+d)",
		C: []Code{
			{0xdd, 0x00},
			{0x45, 0x38},
			{0x00, 0xff},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			r := cpu.regP(codes[1] >> 3)
			p := addrOff(cpu.IX, codes[2])
			*r = cpu.Memory.Get(p)
		},
	},

	OPCode{
		N: "LD r, (IY+d)",
		C: []Code{
			{0xfd, 0x00},
			{0x45, 0x38},
			{0x00, 0xff},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			r := cpu.regP(codes[1] >> 3)
			p := addrOff(cpu.IY, codes[2])
			*r = cpu.Memory.Get(p)
		},
	},

	OPCode{
		N: "LD (HL), r",
		C: []Code{
			{0x70, 0x07},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			r := cpu.regP(codes[0])
			cpu.Memory.Set(cpu.HL.U16(), *r)
		},
	},

	OPCode{
		N: "LD (IX+d), r",
		C: []Code{
			{0xdd, 0x00},
			{0x70, 0x07},
			{0x00, 0xff},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			r := cpu.regP(codes[1])
			p := addrOff(cpu.IX, codes[2])
			cpu.Memory.Set(p, *r)
		},
	},

	OPCode{
		N: "LD (IY+d), r",
		C: []Code{
			{0xfd, 0x00},
			{0x70, 0x07},
			{0x00, 0xff},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			r := cpu.regP(codes[1])
			p := addrOff(cpu.IY, codes[2])
			cpu.Memory.Set(p, *r)
		},
	},

	OPCode{
		N: "LD (HL), n",
		C: []Code{
			{0x36, 0x00},
			{0x00, 0xff},
		},
		T: []int{4, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := cpu.HL.U16()
			cpu.Memory.Set(p, codes[1])
		},
	},

	OPCode{
		N: "LD (IX+d), n",
		C: []Code{
			{0xdd, 0x00},
			{0x36, 0x00},
			{0x00, 0xff},
			{0x00, 0xff},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := addrOff(cpu.IX, codes[2])
			cpu.Memory.Set(p, codes[3])
		},
	},

	OPCode{
		N: "LD (IY+d), r",
		C: []Code{
			{0xfd, 0x00},
			{0x36, 0x00},
			{0x00, 0xff},
			{0x00, 0xff},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := addrOff(cpu.IY, codes[2])
			cpu.Memory.Set(p, codes[3])
		},
	},

	OPCode{
		N: "LD A, (BC)",
		C: []Code{
			{0x0a, 0x00},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := cpu.BC.U16()
			cpu.AF.Hi = cpu.Memory.Get(p)
		},
	},

	OPCode{
		N: "LD A, (DE)",
		C: []Code{
			{0x1a, 0x00},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := cpu.DE.U16()
			cpu.AF.Hi = cpu.Memory.Get(p)
		},
	},

	OPCode{
		N: "LD A, (nn)",
		C: []Code{
			{0x3a, 0x00},
			{0x00, 0xff},
			{0x00, 0xff},
		},
		T: []int{4, 3, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := toU16(codes[1], codes[2])
			cpu.AF.Hi = cpu.Memory.Get(p)
		},
	},

	OPCode{
		N: "LD (BC), A",
		C: []Code{
			{0x02, 0x00},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := cpu.BC.U16()
			cpu.Memory.Set(p, cpu.AF.Hi)
		},
	},

	OPCode{
		N: "LD (DE), A",
		C: []Code{
			{0x12, 0x00},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := cpu.DE.U16()
			cpu.Memory.Set(p, cpu.AF.Hi)
		},
	},

	OPCode{
		N: "LD (nn), A",
		C: []Code{
			{0x32, 0x00},
			{0x00, 0xff},
			{0x00, 0xff},
		},
		T: []int{4, 3, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := toU16(codes[1], codes[2])
			cpu.Memory.Set(p, cpu.AF.Hi)
		},
	},

	OPCode{
		N: "LD A, I",
		C: []Code{
			{0xed, 0x00},
			{0x57, 0x00},
		},
		T: []int{4, 5},
		F: func(cpu *CPU, codes []uint8) {
			d := cpu.IR.Hi
			cpu.AF.Hi = d
			// TODO: update F by d
			// - S is set if the I Register is negative; otherwise, it is
			//   reset.
			// - Z is set if the I Register is 0; otherwise, it is reset.
			// - H is reset.
			// - P/V contains contents of IFF2.
			// - N is reset.
			// - C is not affected.
			// - If an interrupt occurs during execution of this instruction,
			//   the Parity flag contains a 0.
		},
	},

	OPCode{
		N: "LD A, R",
		C: []Code{
			{0xed, 0x00},
			{0x5f, 0x00},
		},
		T: []int{4, 5},
		F: func(cpu *CPU, codes []uint8) {
			d := cpu.IR.Lo
			cpu.AF.Hi = d
			// TODO: update F by d
			// - S is set if, R-Register is negative; otherwise, it is reset.
			// - Z is set if the R Register is 0; otherwise, it is reset.
			// - H is reset.
			// - P/V contains contents of IFF2.
			// - N is reset.
			// - C is not affected.
			// - If an interrupt occurs during execution of this instruction,
			//	 the parity flag contains a 0.
		},
	},

	OPCode{
		N: "LD I, A",
		C: []Code{
			{0xed, 0x00},
			{0x47, 0x00},
		},
		T: []int{4, 5},
		F: func(cpu *CPU, codes []uint8) {
			cpu.IR.Hi = cpu.AF.Hi
		},
	},

	OPCode{
		N: "LD R, A",
		C: []Code{
			{0xed, 0x00},
			{0x4f, 0x00},
		},
		T: []int{4, 5},
		F: func(cpu *CPU, codes []uint8) {
			cpu.IR.Lo = cpu.AF.Hi
		},
	},
}