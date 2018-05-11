package main

import (
	"fmt"
	"strings"
)

func _int_21h() {
	switch REGISTER['m'].(type) {
	case string:
		if strings.Contains(REGISTER['m'].(string), "\\n") {
			lines := removeEmptyStrings(strings.Split(REGISTER['m'].(string), "\\n"))
			debugOutput("Lines: %v\n", lines)
			if len(lines) == 0 {
				fmt.Print("\n")

				return
			}
			for _, val := range lines {
				fmt.Printf("%v\n", strings.Replace(val, "\\n", "", -1))
			}
		} else {
			fmt.Print(REGISTER['m'])
		}
	default:
		fmt.Print(REGISTER['m'])
	}
}
