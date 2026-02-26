package main

import (
	"fmt"
)

func main() {
	// Task 1
	var age int
	age = 26
	fmt.Println(age)
	age = age + 1
	fmt.Println(age)
	// Task 2
	var height int
	height = 180
	fmt.Println(height)
	var height_in_meters float64
	height_in_meters = 1.7
	fmt.Println(height_in_meters)
	// Task 3
	var isStudent bool = true
	fmt.Println(isStudent)
	//task 4
	var temperature int
	fmt.Print("Введите температуру: ")
	fmt.Scanln(&temperature)
	if temperature > 0 {
		fmt.Println("Погода теплая")
	} else {
		fmt.Println("Погода холодная")
	}
	fmt.Println("Task 5")
	var favoriteQuote string
	favoriteQuote = "Запомни, а то забудешь (c)J.Statham,"
	fmt.Println(favoriteQuote)
	// Task 6

	const pi = 3.14
	fmt.Println(pi)
	// pi := 3.1415
	// fmt.Println(pi) // т.к. число является переменным константа, его невозможно менять в последующем коде

}
