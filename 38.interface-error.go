package main
import (
	"errors"
	"fmt"
)


// membuat fungsi dengan penerapan return `error's`
// contoh: func namaFunc(param1, param2) (return1, return2){}
// return 1,2,"etc"

func Pembagian(nilai int, dibagi int)(int, error){
	if dibagi == 0 {
		return 0, errors.New("tdk bisa dibagi dg nol")
	}else{
		return nilai / dibagi, nil
	}
}

func main(){
	var contohErr error = errors.New("test")
	fmt.Println(contohErr.Error())

	fmt.Println(Pembagian(10,2)) // ==> 5 <nil>
	// maka dapat dilakukan iterasi return value kedalam msg2 var

	hasil, err := Pembagian(10,0)
	if err == nil{
		fmt.Println("Hasil :", hasil)
	}else{
		fmt.Println("Error :", err.Error())
	}
}