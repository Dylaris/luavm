package vm

import "luago/api"

/* [A B] R(A) := R(B) */
func move(i Instruction, vm api.LuaVM) {
	// register index
	a, b, _ := i.ABC()
	// convert to stack index
	a += 1
	b += 1
	vm.Copy(b, a)
}

/* [A sBx] PC += sBx */
func jmp(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()
	vm.AddPC(sBx)
	if a != 0 {
		panic("to do")
	}
}
