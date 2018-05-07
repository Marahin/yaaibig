package main

var INSTRUCTION_SET []string

func instructionsCount() int {
	return len(INSTRUCTION_SET)
}

func currentInstruction() int {
	return REGISTER['i']
}

func finishedSourceCode() bool {
	return currentInstruction() == instructionsCount()
}

func getCurrentInstruction() string {
	instruction := INSTRUCTION_SET[currentInstruction()]
	REGISTER['i'] += 1

	return instruction
}
