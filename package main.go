package main

import (
	"fmt"
	"strconv"
)

func main1() {
	/*	var name string
		name = "Alisher"
		fmt.Println(name)

		age:= 25
		println(age)

		name2:="AAA"
		age2:= age
		fmt.Println(name2, age2)

		age = 50
		age2= 55
		fmt.Println(age, age2)

		const pi = 3.14
		fmt.Println(pi) */

	// var schooling int = 11

	// fmt.Println("Studied",schooling)
	// schooling = 12
	// fmt.Println("Studied", schooling)
	// var pi int
	// pi = 3
	// fmt.Println("The value of the pi is", pi)
	// pi2 := float64(pi)
	// fmt.Println("the value of the pi2 is:", pi2)
	// pi2 = 3.14
	// fmt.Println("the updated value of the pi2 is:", pi2)
	a := "10"
	b, _ := strconv.Atoi(a)
	b = b+5
	fmt.Println(b)
}
