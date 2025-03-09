package state

func (ste *luaState) PC() int {
	return ste.pc
}

func (ste *luaState) AddPC(n int) {
	ste.pc += n
}

func (ste *luaState) Fetch() uint32 {
	i := ste.proto.Code[ste.pc]
	ste.pc++
	return i
}

func (ste *luaState) GetConst(idx int) {
	c := ste.proto.Constants[idx]
	ste.stack.push(c)
}

func (ste *luaState) GetRK(rk int) {
	if rk > 0xFF { // constant
		ste.GetConst(rk & 0xFF)
	} else { // register
		ste.PushValue(rk + 1)
	}
}
