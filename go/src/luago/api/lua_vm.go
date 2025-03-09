package api

type LuaVM interface {
	LuaState          // interface embedding
	PC() int          // return current pc
	AddPC(n int)      // add n to the pc
	Fetch() uint32    // fetch current instruction and point to next instruction
	GetConst(idx int) // push the given constant to stack
	GetRK(rk int)     // push given value from constant table or stack to stack
}
