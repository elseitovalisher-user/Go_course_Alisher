package main

import (
	"fmt"
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

}
