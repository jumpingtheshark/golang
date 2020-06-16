package main

import (
	"fmt"
)

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func foo2() {
	e := Employee{
		FirstName:   "Sam",
		LastName:    "Samson",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}

	fmt.Print(e)

}

func main() {
	fmt.Printf("Go Mike!!!\n")

	a := 1
	b := 2
	c := a + b

	fmt.Print(c)
	fmt.Printf("\n")
	fmt.Printf("\n")

	for i := 0; i < 10; i++ {

		fmt.Print(i)
		fmt.Printf("\n")

	}

	foo2()

}
