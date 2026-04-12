package main

import (
	"fmt"
	"math"
)

func main() {

	//task 1
	fmt.Println("Task1")
	var bannerHeight int
	var bannerWidth int
	bannerHeight = 8
	bannerWidth = 12
	bannerArea := bannerHeight * bannerWidth
	fmt.Println(bannerArea)
	var halfBannerArea = bannerArea / 2
	fmt.Println(halfBannerArea)
	var bannerBorderLength = (bannerHeight + bannerWidth) * 2
	fmt.Println(bannerBorderLength)
	//task 2
	fmt.Println("Task2")
	var boxCount = 29
	var numberOfCouriers = 5
	var leftoverBoxes = boxCount % numberOfCouriers
	fmt.Println("leftoverBoxes: ", leftoverBoxes)

	//task 3
	fmt.Println("Task3")
	var tempMorning = 17
	var tempAfternoon = 28
	var tempEvening = 20
	var totalTemp = tempMorning + tempAfternoon + tempEvening
	fmt.Println("Total temp:", totalTemp)
	var averageTemp = (tempMorning + tempAfternoon + tempEvening) / 3
	fmt.Println("average temp:", averageTemp)

	// task 4
	fmt.Println("Task4")
	var knownWords float64 = 47
	var wordsGoal float64 = 120
	var progressPercent = (knownWords / wordsGoal) * 100
	fmt.Println(progressPercent, "%")

	//task 5
	var coins int = 0
	coins += 500
	fmt.Println(coins)
	coins += 1200
	fmt.Println("Bonus:", coins)
	coins /= 2
	fmt.Println("Wasted:", coins)
	coins *= 2
	fmt.Println("Earned x2:", coins)
	coins -= 300
	fmt.Println("Wasted:", coins)
	//task 6
	participants := 42
	groupCount := 8
	participantsPerGroup := participants / groupCount
	fmt.Println(participantsPerGroup)
	// task 7
	fmt.Println(20 - 4*3)
	fmt.Println((20 - 4) * 3)
	/* Так как в программирование, как и в математике сначала выполняются математические действия в скобках,
	а далее приоритет умножение/деление и сложение/вычитание*/

	// task 8
	squareValue := 81

	fmt.Println(math.Sqrt(float64(squareValue)))
	multiplier := 5
	exponent := 3
	fmt.Println(math.Pow(float64(multiplier), (float64(exponent))))

}
