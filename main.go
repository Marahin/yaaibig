package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "error: A path to source code must be passed.\n")
		os.Exit(1)
	}

	inFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: Could not read source from path: %s\nerror: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	defer inFile.Close()
 	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		arguments := strings.Fields(scanner.Text())
		if len(arguments) == 0 {
			continue
		}
		operator_name := arguments[0]
		parameters_stringified := arguments[1:]
		parameters_interfaced := make([]interface{}, len(parameters_stringified))
		for index, val := range parameters_stringified {
			// Try to evalue it into a integer first
			int_value, err := strconv.Atoi(val)
			if err != nil {
				// If that does not work, try to evalue to single rune
				runes := []rune(val)
				if len(runes) > 1 {
					parameters_interfaced[index] = val
				} else {
					parameters_interfaced[index] = runes[0]
				}
			} else {
				parameters_interfaced[index] = int_value
			}
		}
		Call(operator_name, parameters_interfaced...)
  	}

  	gasm_MEMDUMP()
}