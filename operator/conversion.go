package main

import "fmt"

func main() {
	var (
		nil32      = 100000
		nil64      = int64(nil32)
		nil8  int8 = int8(nil32)

		nilIN   int32 = 128
		nilCoba int8  = int8(nilIN)
	)
	fmt.Println(nil32, nil64, nil8, nilCoba)

	// bytecode
	var value1 = "coba"
	// var bytecode = value1[0] //if just 1Arr it'll turn bytecode
	var bytecode = value1[0:4] //if within range it'll turn chop of range
	var valueNum = 10010012
	var bytecodeNum = string(valueNum)
	fmt.Println(bytecode)
	fmt.Println(bytecodeNum)
}
