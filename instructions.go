package main

import "fmt"

var INSTRUCTION_SET []string

func InstructionsCount() int {
	return len(INSTRUCTION_SET)
}

func CurrentInstruction() int {
	return REGISTER['i']
}

func FinishedSourceCode() bool {
	return CurrentInstruction() == InstructionsCount()
}

func GetCurrentInstruction() string {
	instruction := INSTRUCTION_SET[CurrentInstruction()]
	REGISTER['i'] += 1

	return instruction
}

func gasm_INSTRUCTION_DUMP() {
	fmt.Printf("--- gasm_INSTRUCTION_DUMP: ---\n")
	for id, val := range INSTRUCTION_SET {
		fmt.Printf("%v: %s\n", id, val)
	}
}