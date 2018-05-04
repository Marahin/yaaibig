package main

import (
	"fmt"
	"os"
	"reflect"
)

var OPERATORS = map[string]interface{}{
	"MOV": gasm_MOV,
	"mov": gasm_MOV,

	"ADD": gasm_ADD,
	"add": gasm_ADD,

	"MUL": gasm_MUL,
	"mul": gasm_MUL,

	"JNZ": gasm_JNZ,
	"jnz": gasm_JNZ,

	"RET": gasm_RET,
	"ret": gasm_RET,

	"_MEMDUMP": gasm_MEMDUMP,
	"_INSTRUCTION_DUMP": gasm_INSTRUCTION_DUMP,
}

func Call(operator_name string, params ... interface{}) (result []reflect.Value, err error) {
    f := reflect.ValueOf(OPERATORS[operator_name])
    if len(params) != f.Type().NumIn() {
    	fmt.Fprintf(os.Stderr, "error: The number of parameters is not adapted.\n")
        os.Exit(1)
    }
    in := make([]reflect.Value, len(params))
    for k, param := range params {
        in[k] = reflect.ValueOf(param)
    }
    result = f.Call(in)
    return
}

func gasm_MOV(cell rune, value interface{}) {
	switch value_type := value.(type) {
	case rune:
		REGISTER[cell] = REGISTER[value.(rune)]
	case int:
		REGISTER[cell] = value.(int)
    default:
        fmt.Fprintf(os.Stderr, "error: Unsupported type: %T\n", value_type)
        os.Exit(1)
	}
}

func gasm_JNZ(value interface{}) {
	if CurrentMemory() != 0 {
		switch value_type := value.(type) {
		case rune:
			REGISTER['i'] = REGISTER[value.(rune)] - 1
		case int:
			REGISTER['i'] = value.(int) - 1
		default:
	        fmt.Fprintf(os.Stderr, "error: Unsupported type: %T\n", value_type)
	        os.Exit(1)
		}
	}
}

func gasm_ADD(value1, value2 interface{}) {
	sum := 0
	values := []interface{}{value1, value2}
	for _, value := range values {
	    switch value_type := value.(type) {
	    case rune:
	    	sum += REGISTER[value.(rune)]
	    case int:
	        sum += value.(int)
	    default:
	        fmt.Fprintf(os.Stderr, "error: Unsupported type: %T\n", value_type)
	        os.Exit(1)
	    }
	}

	REGISTER['m'] = sum
}

func gasm_RET(value interface{}) {
	switch value_type := value.(type) {
	case rune:
		os.Exit(REGISTER[value.(rune)])
	case int:
		os.Exit(value.(int))
	default:
        fmt.Fprintf(os.Stderr, "error: Unsupported type: %T\n", value_type)
        os.Exit(1)
	}
}

func gasm_MUL(value1, value2 interface{}) {
	base := 0
	multiplier := 0

	switch value1_type := value1.(type) {
	case rune:
		base = REGISTER[value1.(rune)] 
	case int:
		base = value1.(int)
    default:
        fmt.Fprintf(os.Stderr, "error: Unsupported type: %T\n", value1_type)
        os.Exit(1)
	}

	switch value2_type := value2.(type) {
	case rune:
		multiplier = REGISTER[value2.(rune)] 
	case int:
		multiplier = value2.(int)
    default:
        fmt.Fprintf(os.Stderr, "error: Unsupported type: %T\n", value2_type)
        os.Exit(1)
	}

	REGISTER['m'] = base * multiplier
}
