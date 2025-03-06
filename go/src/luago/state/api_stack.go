package state

import (
	"fmt"
	"luago/api"
)

func (ste *luaState) GetTop() int {
	return ste.stack.top
}

func (ste *luaState) AbsIndex(idx int) int {
	return ste.stack.absIndex(idx)
}

func (ste *luaState) CheckStack(n int) bool {
	ste.stack.check(n)
	return true // nerver fails
}

func (ste *luaState) Pop(n int) {
	for range n {
		ste.stack.pop()
	}
}

func (ste *luaState) Copy(fromIdx, toIdx int) {
	val := ste.stack.get(fromIdx)
	ste.stack.set(toIdx, val)
}

func (ste *luaState) PushValue(idx int) {
	val := ste.stack.get(idx)
	ste.stack.push(val)
}

func (ste *luaState) Replace(idx int) {
	val := ste.stack.pop()
	ste.stack.set(idx, val)
}

func (ste *luaState) Insert(idx int) {
	ste.Rotate(idx, 1)
}

func (ste *luaState) Remove(idx int) {
	ste.Rotate(idx, -1)
	ste.Pop(1)
}

func (ste *luaState) Rotate(idx, n int) {
	t := ste.stack.top - 1
	p := ste.stack.absIndex(idx) - 1
	var m int
	if n >= 0 {
		m = t - n
	} else {
		m = p - n - 1
	}
	ste.stack.reverse(p, m)
	ste.stack.reverse(m+1, t)
	ste.stack.reverse(p, t)
}

func (ste *luaState) SetTop(idx int) {
	newTop := ste.stack.absIndex(idx)
	if newTop < 0 {
		panic("stack underflow")
	}

	n := ste.stack.top - newTop
	if n > 0 {
		for range n {
			ste.stack.pop()
		}
	} else if n < 0 {
		for range -n {
			ste.stack.push(nil)
		}
	}
}

func (ste *luaState) PushNil() {
	ste.stack.push(nil)
}

func (ste *luaState) PushBoolean(b bool) {
	ste.stack.push(b)
}

func (ste *luaState) PushInteger(n int64) {
	ste.stack.push(n)
}

func (ste *luaState) PushNumber(n float64) {
	ste.stack.push(n)
}

func (ste *luaState) PushString(s string) {
	ste.stack.push(s)
}

func (ste *luaState) TypeName(tp api.LuaType) string {
	switch tp {
	case api.LUA_TNONE:
		return "no value"
	case api.LUA_TNIL:
		return "nil"
	case api.LUA_TBOOLEAN:
		return "boolean"
	case api.LUA_TNUMBER:
		return "number"
	case api.LUA_TFUNCTION:
		return "function"
	case api.LUA_TTABLE:
		return "table"
	case api.LUA_TSTRING:
		return "string"
	case api.LUA_TTHREAD:
		return "thread"
	default:
		return "userdata"
	}
}

func (ste *luaState) Type(idx int) api.LuaType {
	if ste.stack.isValid(idx) {
		val := ste.stack.get(idx)
		return typeOf(val)
	}
	return api.LUA_TNONE
}

func (ste *luaState) IsNone(idx int) bool {
	return ste.Type(idx) == api.LUA_TNONE
}

func (ste *luaState) IsNil(idx int) bool {
	return ste.Type(idx) == api.LUA_TNIL
}

func (ste *luaState) IsNoneOrNil(idx int) bool {
	t := ste.Type(idx)
	return (t == api.LUA_TNIL || t == api.LUA_TNONE)
}

func (ste *luaState) IsBoolean(idx int) bool {
	return ste.Type(idx) == api.LUA_TBOOLEAN
}

func (ste *luaState) IsInteger(idx int) bool {
	val := ste.stack.get(idx)
	_, ok := val.(int64)
	return ok
}

func (ste *luaState) IsNumber(idx int) bool {
	_, ok := ste.ToNumberX(idx)
	return ok
}

func (ste *luaState) IsString(idx int) bool {
	return ste.Type(idx) == api.LUA_TSTRING
}

func (ste *luaState) ToBoolean(idx int) bool {
	val := ste.stack.get(idx)
	return convertToBoolean(val)
}

func (ste *luaState) ToInteger(idx int) int64 {
	i, _ := ste.ToIntegerX(idx)
	return i
}

func (ste *luaState) ToIntegerX(idx int) (int64, bool) {
	val := ste.stack.get(idx)
	i, ok := val.(int64)
	return i, ok
}

func (ste *luaState) ToNumber(idx int) float64 {
	n, _ := ste.ToNumberX(idx)
	return n
}

func (ste *luaState) ToNumberX(idx int) (float64, bool) {
	val := ste.stack.get(idx)
	switch x := val.(type) {
	case float64:
		return x, true
	case int64:
		return float64(x), true
	default:
		return 0, false
	}
}

func (ste *luaState) ToString(idx int) string {
	s, _ := ste.ToStringX(idx)
	return s
}

func (ste *luaState) ToStringX(idx int) (string, bool) {
	val := ste.stack.get(idx)
	switch x := val.(type) {
	case string:
		return x, true
	case int64, float64:
		s := fmt.Sprintf("%v", x)
		ste.stack.set(idx, s) // modify stack
		return s, true
	default:
		return "", false
	}
}
