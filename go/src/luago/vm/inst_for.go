package vm

import "luago/api"

/* [A sBx] R(A) -= R(A+2); pc += sBx */
func forPrep(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()
	a += 1

	vm.PushValue(a)
	vm.PushValue(a + 2)
	vm.Arith(api.LUA_OPSUB)
	vm.Replace(a)

	vm.AddPC(sBx)
}

/* [A sBx] R(A) += R(A+2); if R(A) <?= R(A+1) then {pc += sBx; R(A+3) = R(A)} */
func forLoop(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()
	a += 1

	vm.PushValue(a)
	vm.PushValue(a + 2)
	vm.Arith(api.LUA_OPADD)
	vm.Replace(a)

	isPositiveStep := vm.ToNumber(a+2) >= 0
	if isPositiveStep && vm.Compare(a, a+1, api.LUA_OPLE) ||
		!isPositiveStep && vm.Compare(a+1, a, api.LUA_OPLE) {
		vm.AddPC(sBx)
		vm.Copy(a, a+3)
	}
}
