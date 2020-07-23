package main

import (
	"fmt"
	"reflect"
	"unsafe"
)
 

 
func main() {
	var data []string // data is bunch of listed string
	for record := 0; record < 1050; record ++{
		data = append(data, fmt.Sprintf("Rec:%d", record)) // output: data => [ "Rec:0", "Rec:1", ...]
		if record < 10 || record == 256 || record == 512 || record == 1024 {

			// &reflect.SliceHeader{} // output: {0, 0, 0} or mean as null struct
			// try to fill `null { 0, 0, 0} struct` with `restricted ptr .(&data)`

			// this : fmt.Printf(("%p\n"),(&data))
			// same kind as: fmt.Println(unsafe.Pointer(&data))

			// below equal as: *reflect.SliceHeader(unsafe.Pointer(&data)) // output : {0, 0, 0}
			// below equal as: *reflect.SliceHeader(0xc0000a0020) // output: {824634523648, 257[record], 512[record]}
			sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
			// fmt.Printf("%p", *sliceHeader)
			fmt.Printf("untuk Index[%d] ->Len[%d] ->Cap[%d]\n",
				record,
				sliceHeader.Len,
				sliceHeader.Cap,
			)
		}
	}
	/*
	var oke int
	oke = (1)
	fmt.Println(oke)
	*/
}
