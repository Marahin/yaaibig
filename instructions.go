package main

var INSTRUCTION_SET []string

func instructionsCount() int {
	return len(INSTRUCTION_SET)
}

func currentInstruction() int {
	return REGISTER['i'].(int)
}

func finishedSourceCode() bool {
	return currentInstruction() == instructionsCount()
}

func getCurrentInstruction() string {
	instruction := INSTRUCTION_SET[currentInstruction()]
	REGISTER['i'] = REGISTER['i'].(int) + 1

	return instruction
}
