package main

import (
	"fmt"
)

func main() {
	// //true
	// fmt.Println(5 == 5)
	// //false
	// fmt.Println(10 != 3)
	// //false
	// fmt.Println(7 > 12)
	// //true
	// fmt.Println(15 < 20)
	// //true
	// fmt.Println(8 >= 8)
	// //false
	// fmt.Println(6 <= 4)
	// //false
	// fmt.Println((10 > 5) && (3 < 1))
	// //true
	// fmt.Println((10 > 5) || (3 < 1))
	// //false
	// fmt.Println(!(5 == 5))
	// //true
	// fmt.Println(!(7 < 3))
	// //false
	// fmt.Println(true && false)
	// //false
	// fmt.Println("tut ", false || false)
	// //true
	// fmt.Println(true || false)
	// //true
	// fmt.Println((4+6 == 10) && (9 > 2))
	// //true
	// fmt.Println((12/3 == 4) || (8 < 5))

	fmt.Println("Task 2")
	var age int
	hasTicket := true
	fmt.Scanln(&age)

	var canEnter bool = age >= 18 && hasTicket
	fmt.Println(canEnter)

	fmt.Println("Task 3")
	var isLoggedIn bool = true
	var isAdmin bool = true
	var hasAccess bool = false
	hasAccess = isLoggedIn && isAdmin || isLoggedIn && !isAdmin

	fmt.Println(hasAccess)
	// fmt.Println(hasAccess && isAdmin)

}
