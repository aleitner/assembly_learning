package main

import (
	"fmt"
)

type VM struct {
	ProgramCtr byte
	Stopped    bool
	BasePtr    int

	Register byte
	Zero     bool
	LessThan bool
	Memory   [256]byte
}

const (
	NOP = iota

	PRINT
	PRINT_POINTER
	PRINT_REGISTER

	STORE_REGISTER_CONSTANT
	SUBTRACT_REGISTER_CONSTANT

	JUMP
	JMP_IF_NOT_ZERO
	JMP_IF_ZERO

	STOP
	ADD

	CMP_REGISTER_CONSTANT

	PUSH_REGISTER
	POP_REGISTER

	CALL
	RETURN
)

func (vm *VM) Step() {
	instruction := vm.Memory[vm.ProgramCtr]
	vm.ProgramCtr++
	switch instruction {
	case NOP:
		break
	case PRINT: // Print value
		arg := vm.Memory[vm.ProgramCtr]
		vm.ProgramCtr++
		fmt.Print(string(arg))
	case PRINT_POINTER: // Print value by pointer
		arg := vm.Memory[vm.ProgramCtr]
		vm.ProgramCtr++

		fmt.Print(vm.Memory[arg])
	case STORE_REGISTER_CONSTANT: // Print value by pointer
		arg := vm.Memory[vm.ProgramCtr]
		vm.ProgramCtr++
		vm.Register = arg
	case PRINT_REGISTER: // Print value in register
		fmt.Print(string(vm.Register))
	case SUBTRACT_REGISTER_CONSTANT:
		arg := vm.Memory[vm.ProgramCtr]
		vm.ProgramCtr++
		vm.Register -= arg

		vm.Zero = vm.Register == 0
	case JUMP:
		arg := vm.Memory[vm.ProgramCtr]
		vm.ProgramCtr = arg
	case JMP_IF_NOT_ZERO:
		arg := vm.Memory[vm.ProgramCtr]
		vm.ProgramCtr++

		if !vm.Zero {
			vm.ProgramCtr = arg
		}
	case JMP_IF_ZERO:
		arg := vm.Memory[vm.ProgramCtr]
		vm.ProgramCtr++

		if vm.Zero {
			vm.ProgramCtr = arg
		}
	case CMP_REGISTER_CONSTANT:
		arg := vm.Memory[vm.ProgramCtr]
		vm.ProgramCtr++

		vm.Zero = arg == vm.Register
	case PUSH_REGISTER:
		vm.Memory[vm.BasePtr] = vm.Register
		vm.BasePtr--
	case POP_REGISTER:
		vm.BasePtr++
		vm.Register = vm.Memory[vm.BasePtr]
	case CALL:
		arg := vm.Memory[vm.ProgramCtr]
		vm.ProgramCtr++

		vm.Memory[vm.BasePtr] = vm.ProgramCtr
		vm.BasePtr--

		vm.ProgramCtr = arg
	case RETURN:
		vm.BasePtr++
		vm.ProgramCtr = vm.Memory[vm.BasePtr]

	case STOP:
		vm.Stopped = true
	default:

	}
}

func main() {
	fmt.Println("Start")

	vm := VM{
		ProgramCtr: 0,
		BasePtr:    255,
		Memory: [256]byte{
			STORE_REGISTER_CONSTANT, 'a',
			CALL, 5,
			STOP,

			PUSH_REGISTER,

			PRINT, 'w',

			STORE_REGISTER_CONSTANT, 6,

			PRINT, 'x',

			SUBTRACT_REGISTER_CONSTANT, 4,

			JMP_IF_NOT_ZERO, 10,

			POP_REGISTER,
			SUBTRACT_REGISTER_CONSTANT, 4,

			JMP_IF_NOT_ZERO, 5,

			RETURN,
		},
	}

	for !vm.Stopped {
		vm.Step()
	}

	fmt.Println("\nEND\n")

}
