package main

import (
	"fmt"
	"sort"
)

var REGISTER = map[rune]int{
	'i': 0, // current instruction registry
	'm': 0, // memory registry
}

func gasm_MEMDUMP() {
	fmt.Printf("--- gasm_MEMDUMP: ---\n")
	var keys []rune
    for k := range REGISTER {
        keys = append(keys, k)
    }
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

    for _, k := range keys {
		fmt.Printf("REGISTER %v = %v\n", string(k), REGISTER[k])
    }
}

func CurrentMemory() int {
	return REGISTER['m']
}