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
	hour := 4
	
	if hour > 0 && hour < 6{
		switch hour {
		fmt.Println("Night")
	}

}
