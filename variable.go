//varuable is place to store data
// each variable is unique
// each identity'es have their own dedicated memory alocation
// can't be relocated each other by their differ new datatypes

// declaring variable name with [0].isNumber is prohibited
// like anyone else language; c/cpp,java,rust,etc
// poc here, uncomment line below
// 1satu_coba := "ini coba"

package main

import "fmt"

func main(){
	var coba string="bisa"
	coba ="lain"
	// cant redeclare used variable
	// even though with difer typeData's
	// poc here, uncomment line below
	// coba = 10
	fmt.Println(coba)


	var ini, untuk, lain string
	fmt.Println(ini, untuk, lain)

	// declare `var` directly without these keyword
	// we can user := whithin 1st variable declaration
	// then if we need to change those variable 'after'
	// we can just filling them with `=` and followed by data
	// example: instead of explicity declaring
	// var like `var ini_data = "isi string"`
	// we can user `ini_data := "isi string"`

	st_data := "re edit me below"
	st_data = "here for editing above"
	fmt.Println(st_data)
}