//varuable is place to store data
// each variable is unique
// each identity'es have their own dedicated memory alocation
// can't be relocated each other by their differ new datatypes

// declaring variable name with [0].isNumber is prohibited
// like anyone else language; c/cpp,java,rust,etc
// poc here, uncomment line below
// 1satu_coba := "ini coba"

// https://golang.org/ref/spec#notation
// place to getting started with
/*
	|   alternation
	()  grouping
	[]  option (0 or 1 times)
	{}  repetition (0 to n times)
*/

/*
The following keywords are reserved and may not be used as identifiers.

break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
The following character sequences represent operators (including assignment operators) and punctuation:
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=
*/

package main

import "fmt"

/*func foo1(parseFunc func(int)int){
	fmt.Println(parseFunc(10))
}
func returnFunc(x int) func(y int) func(){
	return func(y int) func(){
		return func(){
			fmt.Println(x+y)
		}
	}
}*/
func main() {
	/*
		returnFunc(1)(2)()
		// x := returnFunc(2)
		// // fmt.Println(x)
		// x()


		foo1(func(x int) int{
			return x * 11
		})*/

	// =========================================================
	// =========================================================

	// nameVar := Type(option)(another opt) return1 return2{
	// bodies
	// }(directcaller)

	// =========================================================
	// =========================================================

	// nameVar := func(param1 type) returnType{
	// return returnType
	// }(param1)

	// =========================================================
	// =========================================================
	/*

		foo := func(bar int) int{
			// fmt.Println(bar)
			return bar
		}(8) //non-function foo, because foo was called in the end by(8) as param
		fmt.Printf("%T\n",foo)
		fmt.Println(foo)


		foo := func(bar int) int{
			// fmt.Println(bar)
			return bar
		} //function foo, because foo yet called in the end by(8) as param
		foo(9)
		fmt.Printf("%T\n",foo)
		fmt.Println(foo(10)) //only recently used at this point
	*/
	var (
		var1 bool   = false
		var2 bool   = true
		var3 string = "tiga"
	)
	fmt.Println(var1, var2, var3)
	fmt.Println("\n\no\r\rpp")

	var coba string = "bisa"
	coba = "lain"
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

	// variable multiple declaration
	// it's use an array
	// an_obj(arg_ar1, arg_ar2, arg_etc)
	// poc here, uncomment line below
	// var(
	// 	my1stName = "eko"
	// 	my2ndName = "mulyadi"
	// )
	var (
		my1stName = "eko"
		my2ndName = "mulyadi"
	)
	fmt.Println(my1stName, my2ndName)
}
