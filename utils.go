package main

import (
	"fmt"
	"flag"
)

var DEBUG_FLAG bool

func removeEmptyStrings(arr []string) []string {
	var newArr = make([]string, 0)

	for _, val := range arr {
		if val != "" {
			newArr = append(newArr, val)
		}
	}

	return newArr
}

func parseCLF() {
	flag.BoolVar(&DEBUG_FLAG, "debug", false, "adds a very verbose output")
	flag.Parse()
}

func debugOutput(format string, params ...interface{}) {
	if DEBUG_FLAG {
		fmt.Printf(format, params...)
	}
}