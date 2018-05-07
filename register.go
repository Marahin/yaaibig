package main

var REGISTER = map[rune]int{
	'i': 0, // current instruction registry
	'm': 0, // memory registry
}

func currentMemory() int {
	return REGISTER['m']
}
