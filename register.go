package main

import (
	"fmt"
	"sort"
)

var REGISTER = make(map[rune]int)

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