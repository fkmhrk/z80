package z80

import "math/bits"

func (cpu *CPU) addU8(a, b uint8) uint8 {
	v := uint16(a) + uint16(b)
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v&0xff == 0).
		Put(H, a&0x0f+b&0x0f > 0x0f).
		// TODO: verify PV behavior.
		Put(PV, a&0x80 == b&0x80 && a&0x80 != uint8(v&0x80)).
		Reset(N).
		Put(C, v > 0xff))
	return uint8(v)
}

func (cpu *CPU) adcU8(a, b uint8) uint8 {
	a16 := uint16(a)
	b16 := uint16(b)
	if cpu.flag(C) {
		b16++
	}
	v := a16 + b16
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v&0xff == 0).
		// TODO: verify H behavior.
		Put(H, a&0x0f+b&0x0f > 0x0f).
		// TODO: verify PV behavior.
		Put(PV, a&0x80 == b&0x80 && a&0x80 != uint8(v&0x80)).
		Reset(N).
		Put(C, v > 0xff))
	return uint8(v)
}

func (cpu *CPU) subU8(a, b uint8) uint8 {
	v := uint16(a) - uint16(b)
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v&0xff == 0).
		Put(H, a&0x0f < b&0x0f).
		// TODO: verify PV behavior.
		Put(PV, a&0x80 == b&0x80 && a&0x80 != uint8(v&0x80)).
		Set(N).
		Put(C, v > 0xff))
	return uint8(v)
}

func (cpu *CPU) sbcU8(a, b uint8) uint8 {
	a16 := uint16(a)
	b16 := uint16(b)
	if cpu.flag(C) {
		b16++
	}
	v := a16 - b16
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v&0xff == 0).
		// TODO: verify H behavior.
		Put(H, a&0x0f < b&0x0f).
		// TODO: verify PV behavior.
		Put(PV, a&0x80 == b&0x80 && a&0x80 != uint8(v&0x80)).
		Set(N).
		Put(C, v > 0xff))
	return uint8(v)
}

func (cpu *CPU) andU8(a, b uint8) uint8 {
	v := a & b
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Set(H).
		// TODO: verify PV behavior.
		Put(PV, bits.OnesCount8(v)%2 == 0).
		Reset(N).
		Reset(C))
	return uint8(v)
}

func (cpu *CPU) orU8(a, b uint8) uint8 {
	v := a | b
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Reset(H).
		// TODO: verify PV behavior.
		Put(PV, bits.OnesCount8(v)%2 == 0).
		Reset(N).
		Reset(C))
	return uint8(v)
}

func (cpu *CPU) xorU8(a, b uint8) uint8 {
	v := a ^ b
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Reset(H).
		// TODO: verify PV behavior.
		Put(PV, bits.OnesCount8(v)%2 == 0).
		Reset(N).
		Reset(C))
	return uint8(v)
}

func (cpu *CPU) incU8(a uint8) uint8 {
	v := a + 1
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Put(H, a&0x0f+1 > 0x0f).
		Put(PV, a == 0x7f).
		Reset(N))
	return v
}

func (cpu *CPU) decU8(a uint8) uint8 {
	v := a - 1
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Put(H, a&0x0f < 1).
		Put(PV, a == 0x80).
		Reset(N))
	return v
}

func (cpu *CPU) addU16(a, b uint16) uint16 {
	v := uint32(a) + uint32(b)
	cpu.flagUpdate(FlagOp{}.
		Put(H, a&0x0fff+b&0x0fff > 0x0fff).
		Reset(N).
		Put(C, v > 0xffff))
	return uint16(v)
}

func (cpu *CPU) adcU16(a, b uint16) uint16 {
	a16 := uint32(a)
	b16 := uint32(a)
	if cpu.flag(C) {
		b16++
	}
	v := a16 + b16
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x8000 != 0).
		Put(Z, v&0xffff == 0).
		// TODO: verify H behavior.
		Put(H, a&0x0fff+b&0x0fff > 0x0fff).
		// TODO: verify PV behavior.
		Put(PV, a&0x8000 == b&0x8000 && a&0x8000 != uint16(v&0x8000)).
		Reset(N).
		Put(C, v > 0xffff))
	return uint16(v)
}

func (cpu *CPU) sbcU16(a, b uint16) uint16 {
	a16 := int32(a)
	b16 := int32(a)
	if cpu.flag(C) {
		b16++
	}
	v := a16 - b16
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x8000 != 0).
		Put(Z, v&0xffff == 0).
		// TODO: verify H behavior.
		Put(H, a&0x0fff < b&0x0fff).
		// TODO: verify PV behavior.
		Put(PV, a&0x8000 == b&0x8000 && a&0x8000 != uint16(v&0x8000)).
		Reset(N).
		Put(C, v < 0x0000))
	return uint16(v)
}

func (cpu *CPU) incU16(a uint16) uint16 {
	v := a + 1
	return v
}

func (cpu *CPU) decU16(a uint16) uint16 {
	v := a - 1
	return v
}

func (cpu *CPU) rlcU8(a uint8) uint8 {
	a2 := a<<1 | a>>7
	cpu.flagUpdate(FlagOp{}.
		Put(S, a2&0x80 != 0).
		Put(Z, a2 == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(a2)%2 == 0).
		Reset(N).
		Put(C, a&0x80 != 0))
	return a2
}

func (cpu *CPU) rlU8(a uint8) uint8 {
	a2 := a << 1
	if cpu.flag(C) {
		a2 |= 0x01
	}
	cpu.flagUpdate(FlagOp{}.
		Put(S, a2&0x80 != 0).
		Put(Z, a2 == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(a2)%2 == 0).
		Reset(N).
		Put(C, a&0x80 != 0))
	return a2
}

func (cpu *CPU) rrcU8(a uint8) uint8 {
	a2 := a>>1 | a<<7
	cpu.flagUpdate(FlagOp{}.
		Put(S, a2&0x80 != 0).
		Put(Z, a2 == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(a2)%2 == 0).
		Reset(N).
		Put(C, a&0x01 != 0))
	return a2
}

func (cpu *CPU) rrU8(a uint8) uint8 {
	a2 := a >> 1
	if cpu.flag(C) {
		a2 |= 0x80
	}
	cpu.flagUpdate(FlagOp{}.
		Put(S, a2&0x80 != 0).
		Put(Z, a2 == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(a2)%2 == 0).
		Reset(N).
		Put(C, a&0x01 != 0))
	return a2
}

func (cpu *CPU) slaU8(a uint8) uint8 {
	a2 := a << 1
	cpu.flagUpdate(FlagOp{}.
		Put(S, a2&0x80 != 0).
		Put(Z, a2 == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(a2)%2 == 0).
		Reset(N).
		Put(C, a&0x80 != 0))
	return a2
}

func (cpu *CPU) sraU8(a uint8) uint8 {
	a2 := a&0x80 | a>>1
	cpu.flagUpdate(FlagOp{}.
		Put(S, a2&0x80 != 0).
		Put(Z, a2 == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(a2)%2 == 0).
		Reset(N).
		Put(C, a&0x01 != 0))
	return a2
}

func (cpu *CPU) srlU8(a uint8) uint8 {
	a2 := a >> 1
	cpu.flagUpdate(FlagOp{}.
		Put(S, a2&0x80 != 0).
		Put(Z, a2 == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(a2)%2 == 0).
		Reset(N).
		Put(C, a&0x01 != 0))
	return a2
}
