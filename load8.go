package z80

var load8 = []*OPCode{

	{
		N: "LD r1, r2",
		C: []Code{
			{0x40, 0x3f, vReg88},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			*cpu.regP(codes[0] >> 3) = *cpu.regP(codes[0])
		},
	},

	{
		N: "LD r, n",
		C: []Code{
			{0x06, 0x38, vReg88},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			r := cpu.regP(codes[0] >> 3)
			*r = codes[1]
		},
	},

	{
		N: "LD r, (HL)",
		C: []Code{
			{0x45, 0x38, vReg8_3},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			r := cpu.regP(codes[0] >> 3)
			*r = cpu.Memory.Get(cpu.HL.U16())
		},
	},

	{
		N: "LD r, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x45, 0x38, vReg8_3},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			r := cpu.regP(codes[1] >> 3)
			p := addrOff(cpu.IX, codes[2])
			*r = cpu.Memory.Get(p)
		},
	},

	{
		N: "LD r, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x45, 0x38, vReg8_3},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			r := cpu.regP(codes[1] >> 3)
			p := addrOff(cpu.IY, codes[2])
			*r = cpu.Memory.Get(p)
		},
	},

	{
		N: "LD (HL), r",
		C: []Code{
			{0x70, 0x07, vReg8},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			r := cpu.regP(codes[0])
			cpu.Memory.Set(cpu.HL.U16(), *r)
		},
	},

	{
		N: "LD (IX+d), r",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x70, 0x07, vReg8},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			r := cpu.regP(codes[1])
			p := addrOff(cpu.IX, codes[2])
			cpu.Memory.Set(p, *r)
		},
	},

	{
		N: "LD (IY+d), r",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x70, 0x07, vReg8},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			r := cpu.regP(codes[1])
			p := addrOff(cpu.IY, codes[2])
			cpu.Memory.Set(p, *r)
		},
	},

	{
		N: "LD (HL), n",
		C: []Code{
			{0x36, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := cpu.HL.U16()
			cpu.Memory.Set(p, codes[1])
		},
	},

	{
		N: "LD (IX+d), n",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x36, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := addrOff(cpu.IX, codes[2])
			cpu.Memory.Set(p, codes[3])
		},
	},

	{
		N: "LD (IY+d), r",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x36, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := addrOff(cpu.IY, codes[2])
			cpu.Memory.Set(p, codes[3])
		},
	},

	{
		N: "LD A, (BC)",
		C: []Code{
			{0x0a, 0x00, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := cpu.BC.U16()
			cpu.AF.Hi = cpu.Memory.Get(p)
		},
	},

	{
		N: "LD A, (DE)",
		C: []Code{
			{0x1a, 0x00, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := cpu.DE.U16()
			cpu.AF.Hi = cpu.Memory.Get(p)
		},
	},

	{
		N: "LD A, (nn)",
		C: []Code{
			{0x3a, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := toU16(codes[1], codes[2])
			cpu.AF.Hi = cpu.Memory.Get(p)
		},
	},

	{
		N: "LD (BC), A",
		C: []Code{
			{0x02, 0x00, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := cpu.BC.U16()
			cpu.Memory.Set(p, cpu.AF.Hi)
		},
	},

	{
		N: "LD (DE), A",
		C: []Code{
			{0x12, 0x00, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := cpu.DE.U16()
			cpu.Memory.Set(p, cpu.AF.Hi)
		},
	},

	{
		N: "LD (nn), A",
		C: []Code{
			{0x32, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := toU16(codes[1], codes[2])
			cpu.Memory.Set(p, cpu.AF.Hi)
		},
	},

	{
		N: "LD A, I",
		C: []Code{
			{0xed, 0x00, nil},
			{0x57, 0x00, nil},
		},
		T: []int{4, 5},
		F: func(cpu *CPU, codes []uint8) {
			d := cpu.IR.Hi
			cpu.AF.Hi = d
			// update F by d
			// - S is set if the I Register is negative; otherwise, it is
			//   reset.
			// - Z is set if the I Register is 0; otherwise, it is reset.
			// - H is reset.
			// - P/V contains contents of IFF2.
			// - N is reset.
			// - C is not affected.
			// - If an interrupt occurs during execution of this instruction,
			//   the Parity flag contains a 0.
			cpu.flagUpdate(FlagOp{}.
				Put(S, d&0x80 != 0).
				Put(Z, d == 0).
				Reset(H).
				Put(PV, cpu.IFF2).
				Reset(N).
				Keep(C))
		},
	},

	{
		N: "LD A, R",
		C: []Code{
			{0xed, 0x00, nil},
			{0x5f, 0x00, nil},
		},
		T: []int{4, 5},
		F: func(cpu *CPU, codes []uint8) {
			d := cpu.IR.Lo
			cpu.AF.Hi = d
			// update F by d
			// - S is set if, R-Register is negative; otherwise, it is reset.
			// - Z is set if the R Register is 0; otherwise, it is reset.
			// - H is reset.
			// - P/V contains contents of IFF2.
			// - N is reset.
			// - C is not affected.
			// - If an interrupt occurs during execution of this instruction,
			//	 the parity flag contains a 0.
			cpu.flagUpdate(FlagOp{}.
				Put(S, d&0x80 != 0).
				Put(Z, d == 0).
				Reset(H).
				Put(PV, cpu.IFF2).
				Reset(N).
				Keep(C))
		},
	},

	{
		N: "LD I, A",
		C: []Code{
			{0xed, 0x00, nil},
			{0x47, 0x00, nil},
		},
		T: []int{4, 5},
		F: func(cpu *CPU, codes []uint8) {
			cpu.IR.Hi = cpu.AF.Hi
		},
	},

	{
		N: "LD R, A",
		C: []Code{
			{0xed, 0x00, nil},
			{0x4f, 0x00, nil},
		},
		T: []int{4, 5},
		F: func(cpu *CPU, codes []uint8) {
			cpu.IR.Lo = cpu.AF.Hi
		},
	},
}
