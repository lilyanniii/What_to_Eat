package main

import "fmt"

func main() {
	fmt.Println("What type of food did you want to eat?")
	fmt.Println("\nYour choices are American, Chinese, Fast Food, and your Favorites")

	typeOfFood1 := "American"
	typeOfFood2 := "Chinese"
	typeOfFood3 := "FastFood"
	typeOfFood4 := "Favorites"

	//Learn how to randomize these lists
	American := "\nMarty's Grill \nRuby Tuesday \nRiverbound Cafe \nCracker Barrel \nSports Page Bar \nTGI Fridays \nCharred Hanover \nOutback Steakhouse \nWaffle House"
	Chinese := "\nChen's Chinese \nBell Hut Asian Cuisine  \nZheng Chinese Restaurant \nChina Kitchen \nChina Wok \nPeking Restaurant \nChina House \nGinger Red Asian Bistro \nTops China \nFuji Asian Cuisine \nMarigold Thai \nBonchon \nLittle Szechuan"
	fastFood := "\nChick-fil-A \nWendy's \nArby's \nCook Out \nMcDonald's \nTaco Bell \nSonic \nFive Guys \nBurger King \nSubway"
	Favorites := "\nZaxby's \nChic-fil-A"

	var typeOfFood string

	fmt.Scan(&typeOfFood)

	switch typeOfFood {
	case typeOfFood1:
		fmt.Println(American)
	case typeOfFood2:
		fmt.Println(Chinese)
	case typeOfFood3:
		fmt.Println(fastFood)
	case typeOfFood4:
		fmt.Println(Favorites)
	}
}
