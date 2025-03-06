package state

type luaStack struct {
	slots []luaValue
	top   int
}

func newLuaStack(size int) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top:   0,
	}
}

func (stk *luaStack) check(n int) {
	free := len(stk.slots) - stk.top
	for i := free; i < n; i++ {
		stk.slots = append(stk.slots, nil)
	}
}

func (stk *luaStack) push(val luaValue) {
	if stk.top == len(stk.slots) {
		panic("stack overflow")
	}
	stk.slots[stk.top] = val
	stk.top++
}

func (stk *luaStack) pop() luaValue {
	if stk.top < 1 {
		panic("stack underflow")
	}
	stk.top--
	val := stk.slots[stk.top]
	stk.slots[stk.top] = nil
	return val
}

func (stk *luaStack) absIndex(idx int) int {
	if idx >= 0 {
		return idx
	}
	return idx + stk.top + 1
}

func (stk *luaStack) isValid(idx int) bool {
	absIdx := stk.absIndex(idx)
	return absIdx > 0 && absIdx <= stk.top
}

func (stk *luaStack) get(idx int) luaValue {
	absIdx := stk.absIndex(idx)
	if absIdx > 0 && absIdx <= stk.top {
		return stk.slots[absIdx-1]
	}
	return nil
}

func (stk *luaStack) set(idx int, val luaValue) {
	absIdx := stk.absIndex(idx)
	if absIdx > 0 && absIdx <= stk.top {
		stk.slots[absIdx-1] = val
		return
	}
	panic("invalid index")
}

func (stk *luaStack) reverse(from, to int) {
	slots := stk.slots
	for from < to {
		slots[from], slots[to] = slots[to], slots[from]
		from++
		to--
	}
}
