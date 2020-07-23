package main

import "fmt"

type Salary struct{
	Basic, HRA, TA float64
}

type Employee struct{
	FirstName, Lastname, Email string
	Age           int
	MonthlySalary []Salary // put Salary struct as slice of arr
}

// set e as param-get from Employee struct
func (e Employee) EmpInfo() string{ //must explicity return a `string` that warped by EmpInfo()

	//return any string
	fmt.Println(e.FirstName, e.Lastname)
	fmt.Println(e.Age)
	fmt.Println(e.Email)

	// _ => denoted as index of slice
	// iterate _ and info from range slice of `e.MonthlySalary` array
	for i, info := range e.MonthlySalary{
		fmt.Printf("===[ BAYARAN %d ]===\n",i)
		fmt.Println(info.Basic)
		fmt.Println(info.HRA)
		fmt.Println(info.TA)
	}
	return "-----------------"
}

func main(){
	e := Employee{
		FirstName: "Ahmad",
		Lastname: "Mutadho",
		Email: "Ammumat@gmail.com",
		Age: 29,
		MonthlySalary: []Salary{
			Salary{
				Basic: 20.000,
				HRA: 19.000,
				TA: 80.000,
			},
			Salary{
				Basic: 18.000,
				HRA: 11.000,
				TA: 40.000,
			},
		},
	}
	fmt.Println(e.EmpInfo())
}
