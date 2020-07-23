package main

import "fmt"

func main(){
	if 7%5==0{ //if 7 is divisible by 5
		fmt.Println("buset bg")
	} else{ //penulisan else di golang harus 1line setelah tutup if{}
		fmt.Println("kayae tidak mungkin. males mau jabarin")
	}

	if(10+10==20) {
		// check color formating...
		// buset..
		// digolang statment + () malah jadi kek fungsi :/
		// meski statement tetep worksih
		fmt.Println("adalah benar")
	}
	if(10+10)==20{
		fmt.Println("juga sama benarnya")
	}
	if false{
		fmt.Println("lanjut else if bawahnya yaa")
	}else if(true){
		fmt.Println("mantap ketemu dengan saya disini")
	}else{
		fmt.Println("tidak akan pernah bertemu dengan saya")
	}
	
	if num:=9; num < 0{ //mirip dengan requirement and
		fmt.Println("tidak akan muncul karena salah")
	}
	if num:=9; num == 9{
		fmt.Println("alternative. dan akan muncul karena benar")
	}

}