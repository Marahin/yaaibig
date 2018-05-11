package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	parseCLF()

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "error: A path to source code must be passed.\n")
		os.Exit(1)
	}

	// Ignore flags
	argumentsStartIndex := 1
	for id, val := range os.Args[1:] {
		if string(val[0]) != "-" {
			argumentsStartIndex += id // += because we start from os.Args[1]
		}
	}
	filePath := os.Args[argumentsStartIndex]

	inFile, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: Could not read source from path: %s\nerror: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		rawLine := scanner.Text()
		INSTRUCTION_SET = append(INSTRUCTION_SET, rawLine)
	}

	for !finishedSourceCode() {
		rawLine := getCurrentInstruction()

		operatorName, parameters, err := parseLine(rawLine)
		debugOutput("parameters: %v\nlen(parameters):%v\n", parameters, len(parameters))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: Parsing error on line\nerror: %v\n", REGISTER['i'].(int), err)
			os.Exit(1)
		}

		if operatorName == "" && len(parameters) == 0 { // Nothing to do!
			continue
		}

		call(operatorName, parameters...)
	}
}
