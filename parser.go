package main

import (
	"errors"
	"strings"
	"strconv"
)

func parseLine(line string) (string, []interface{}, error) {
	// Get line without any comments (ignored)
	lineWoComments := strings.Split(line, ";")[0]

	if len(lineWoComments) == 0 { // No code
		return "", nil, nil
	}

	// Catch all string arguments first 
	// (MOV A "test") == []string{"MOV A", "test"}
	argumentsAndStringArguments := removeEmptyStrings(strings.Split(lineWoComments, "\""))
	nonStringArguments := argumentsAndStringArguments[0]
	var arguments = make([]string, 0)

	if len(argumentsAndStringArguments) > 1 { // Received string argument
		stringArguments := argumentsAndStringArguments[1:]

		arguments = strings.Fields(nonStringArguments)
		arguments = append(arguments, stringArguments...)
	} else {
		arguments = strings.Fields(nonStringArguments)
	}

	if len(arguments) == 0 {
		return "", make([]interface{}, 0), errors.New("No arguments passed")
	}

	debugOutput("arguments: %v\n", arguments)

	operatorName := arguments[0]
	debugOutput("operatorName: %v\n", operatorName)
	parametersStringified := arguments[1:]
	debugOutput("parametersStringified: %v\nlen(parametersStringified): %v\n", parametersStringified, len(parametersStringified))
	parametersInterfaced := make([]interface{}, len(parametersStringified))
	debugOutput("parametersInterfaced: %v\nlen(parametersInterfaced): %v\n", parametersInterfaced, len(parametersInterfaced))

	for id, val := range parametersStringified {
		debugOutput("val:%v\n", val)
		// Catch integers first
		intVal, err := strconv.Atoi(val)
		if err == nil { // Catched an integer
			parametersInterfaced[id] = intVal
			debugOutput("Found int: (val: %v, intVal: %v)\n", val, intVal)
			continue
		} // Loop exited if integer catched

		// Catch runes
		runes := []rune(val)
		if len(runes) > 1 {
			parametersInterfaced[id] = val // String catched
			continue
		}

		parametersInterfaced[id] = runes[0]	 // Rune catched
	}


	return operatorName, parametersInterfaced, nil
}

