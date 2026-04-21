// Homework 10 - Условные конструкции

package main

import (
	"fmt"
)

func main() {
	//task 1
	temperature := -20
	if temperature < 0 {
		fmt.Println("Холодно")
	} else if temperature >= 0 && temperature <= 20 {
		fmt.Println("Тепло")
	} else {
		fmt.Println("Жарко")
	}
	// task 2
	score := 80
	if score >= 90 && score == 100 {
		fmt.Println("Excellent")
	} else if score >= 70 && score <= 89 {
		fmt.Println("Good")
	} else if score > 50 && score < 70 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
	// task 3
	hour := 26
	switch {

	case hour >= 6 && hour <= 11:
		fmt.Println("Morning")
	case hour >= 12 && hour <= 17:
		fmt.Println("Day")
	case hour >= 18 && hour <= 23:
		fmt.Println("Evening")
	}
	switch hour {
	case 0, 1, 2, 3, 4, 5:
		fmt.Println("Night")
	default:
		fmt.Println("Out of clock range")
	}
	// Task 4
	var number int = -10000

	fmt.Scanln(&number)
	if number == -10000 {
		fmt.Println("Введено неверное число")
	} else if number%2 == 0 {
		fmt.Println("Четное число")
	} else {
		fmt.Println("Нечетное число")
	}
	// task 5
	day1 := "Monday"
	day2 := "Tuesday"
	day3 := "Wednsday"
	day4 := "Thursday"
	day5 := "Friday"
	day6 := "Saturday"
	day7 := "Sunday"
	day := day4

	switch day {
	case day1, day2, day3, day4, day5:
		fmt.Println("Work Day")
	case day6, day7:
		fmt.Println("Weekend")

	default:
		fmt.Println("Not a weekday")
	}
	// task 6
	fmt.Println("Task 5")
	balance := 0
	if balance > 0 {
		fmt.Println("Баланс положительный")
	} else {
		fmt.Println("Баланс отрицательный")
	}
	fmt.Println("Task 7")
	//task 7
	var age int
	fmt.Scanln(&age)
	if age < 13 && age > 0 {
		fmt.Println("Ребенок")
	} else if age >= 13 && age < 17 {
		fmt.Println("ПОдросток")
	} else if age > 18 {
		fmt.Println("Взрослый")
	} else {
		fmt.Println("Неверный возраст")
	}
	//task 8
	var command string
	fmt.Scanln(&command)
	switch command {
	case "start":
		fmt.Println("Programm is starting...")
	case "stop":
		fmt.Println("Programm is stopped...")
	case "restart":
		fmt.Println("Programm is restarting...")

	default:
		fmt.Println("Unknown command")

	}
	//task 9
	var grade int
	fmt.Scanln(&grade)
	switch grade {
	case 1:
		fmt.Println("F")
	case 2:
		fmt.Println("D")
	case 3:
		fmt.Println("C")
	case 4:
		fmt.Println("B")
	case 5:
		fmt.Println("A")
	default:
		fmt.Println("Некорректная оценка")
	}

}
