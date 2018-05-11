package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
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

	"JMP": gasm_JMP,
	"jmp": gasm_JMP,

	"RET": gasm_RET,
	"ret": gasm_RET,

	"INT": gasm_INT,
	"int": gasm_INT,

	"_MEMDUMP":          gasm_MEMDUMP,
	"_INSTRUCTION_DUMP": gasm_INSTRUCTION_DUMP,
}

func call(operator_name string, params ...interface{}) (result []reflect.Value, err error) {
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
	case string:
		REGISTER[cell] = value.(string)
	default:
		fmt.Fprintf(os.Stderr, "error: Unsupported type: %T\n", value_type)
		os.Exit(1)
	}
}

func gasm_INT(value interface{}) {
	debugOutput("REGISTER['m']: %v\n", REGISTER['m'])
	switch value_type := value.(type) {
	case string:
		switch value.(string) {
		case "21h":
			_int_21h()
		}
	default:
		fmt.Fprintf(os.Stderr, "error: Unsupported type: %T\n", value_type)
		os.Exit(1)
	}
}

func gasm_JNZ(value interface{}) {
	if currentMemory() != 0 {
		switch value_type := value.(type) {
		case rune:
			REGISTER['i'] = REGISTER[value.(rune)].(int) - 1
		case int:
			REGISTER['i'] = value.(int) - 1
		default:
			fmt.Fprintf(os.Stderr, "error: Unsupported type: %T\n", value_type)
			os.Exit(1)
		}
	}
}

func gasm_JMP(value interface{}) {
	switch value_type := value.(type) {
	case rune:
		REGISTER['i'] = REGISTER[value.(rune)].(int) - 1
	case int:
		REGISTER['i'] = value.(int) - 1
	default:
		fmt.Fprintf(os.Stderr, "error: Unsupported type: %T\n", value_type)
		os.Exit(1)
	}
}

func gasm_ADD(value1, value2 interface{}) {
	sum := 0
	values := []interface{}{value1, value2}
	for _, value := range values {
		switch value_type := value.(type) {
		case rune:
			sum += REGISTER[value.(rune)].(int)
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
		os.Exit(REGISTER[value.(rune)].(int))
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
		base = REGISTER[value1.(rune)].(int)
	case int:
		base = value1.(int)
	default:
		fmt.Fprintf(os.Stderr, "error: Unsupported type: %T\n", value1_type)
		os.Exit(1)
	}

	switch value2_type := value2.(type) {
	case rune:
		multiplier = REGISTER[value2.(rune)].(int)
	case int:
		multiplier = value2.(int)
	default:
		fmt.Fprintf(os.Stderr, "error: Unsupported type: %T\n", value2_type)
		os.Exit(1)
	}

	REGISTER['m'] = base * multiplier
}

func gasm_INSTRUCTION_DUMP() {
	fmt.Printf("--- gasm_INSTRUCTION_DUMP: ---\n")
	for id, val := range INSTRUCTION_SET {
		fmt.Printf("%v: %s\n", id, val)
	}
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
