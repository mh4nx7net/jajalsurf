package main


import "fmt"

func main(){
	var coba int
	for(coba=0; coba<10;coba++){
		fmt.Println("cok")
	}
	if true && true {
		fmt.Println("and logic")
	}
	if true && false {
		fmt.Println("and logic; no printout")
	}
	if false || true {
		fmt.Println("or logic")
	}
	if false || false {
		fmt.Println("or logic; no printout")
	}
}
