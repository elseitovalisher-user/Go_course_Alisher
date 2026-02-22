package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")

	//Task 1
	var schooling int = 11

	fmt.Println("Studied", schooling)
	schooling = 12
	fmt.Println("Studied", schooling)

	fmt.Println("Task 2")	
	var name string
	name = "Vladislav"
	fmt.Print(name)
	name = " Alisher"
	fmt.Println(name)

	fmt.Println("Task 3")
	var steps int
	steps = 0
	fmt.Println("morning steps: ", steps)
	steps = 2000
	fmt.Println("afternoon steps: ", steps)
	fmt.Println("Хорошая работа! Вы уже на пути к своей ежедневной цели\n",
		"Well done")
	var largeNumber int
	largeNumber = 5000055
	fmt.Println("Task 4", largeNumber)

	const breakTime = 15
		fmt.Println("task 5", breakTime)
	//const breakTime = 20
	//fmt.Println(breakTime) // т.к. число является переменным константа, его невозможно менять в последующем коде

}
