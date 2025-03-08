package state

func (ste *luaState) Len(idx int) {
	val := ste.stack.get(idx)
	if s, ok := val.(string); ok {
		ste.stack.push(int64(len(s)))
	} else {
		panic("length error")
	}
}

func (ste *luaState) Concat(n int) {
	if n == 0 {
		ste.stack.push("")
	} else if n >= 1 {
		for range n - 1 {
			if ste.IsString(-1) && ste.IsString(-2) {
				s2 := ste.ToString(-1)
				s1 := ste.ToString(-2)
				ste.stack.pop()
				ste.stack.pop()
				ste.stack.push(s1 + s2)
				continue
			}
			panic("concatentation error")
		}
	}
	// n == 1 do nothing
}
