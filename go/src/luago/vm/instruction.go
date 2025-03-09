package vm

import "luago/api"

/* Decoding instructions */

type Instruction uint32

const MAXARG_Bx = 1<<18 - 1
const MAXARG_sBx = MAXARG_Bx >> 1

func (inst Instruction) Opcode() int {
	return int(inst & 0x3F)
}

func (inst Instruction) ABC() (a, b, c int) {
	a = int(inst >> 6 & 0xFF)
	c = int(inst >> 14 & 0x1FF)
	b = int(inst >> 23 & 0x1FF)
	return
}

func (inst Instruction) ABx() (a, bx int) {
	a = int(inst >> 6 & 0xFF)
	bx = int(inst >> 14)
	return
}

func (inst Instruction) AsBx() (a, sbx int) {
	a, bx := inst.ABx()
	return a, bx - MAXARG_sBx
}

func (inst Instruction) Ax() int {
	return int(inst >> 6)
}

func (inst Instruction) OpName() string {
	return opcodes[inst.Opcode()].name
}

func (inst Instruction) OpMode() byte {
	return opcodes[inst.Opcode()].opMode
}

func (inst Instruction) BMode() byte {
	return opcodes[inst.Opcode()].argBMode
}

func (inst Instruction) CMode() byte {
	return opcodes[inst.Opcode()].argCMode
}

func (inst Instruction) Execute(vm api.LuaVM) {
	action := opcodes[inst.Opcode()].action
	if action != nil {
		action(inst, vm)
	} else {
		panic(inst.OpName())
	}
}
