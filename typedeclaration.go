package main
import "fmt"

func main(){
	// `type`	is for aliases type of data's
	// poc here, example below
	// type byte = uint8
	// type rune = int32
	// type uint = uint

	type cek bool
	var adalahBenar cek = true
	// type salah false
	if(adalahBenar){
		fmt.Println("benar")
	}

	type married = bool
	isMarried := true
	if isMarried{
		fmt.Println("sudah rabi")
	}
	
}