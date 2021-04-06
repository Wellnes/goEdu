package main

import "fmt"

func main() {

	// var price int
	// fmt.Scan(&price)

	// switch price {
	// case 100:
	// 	fmt.Println("First case")
	// case 110:
	// 	fmt.Println("Second case")
	// case 120:
	// 	fmt.Println("Third case")
	// default:
	// 	fmt.Println("Default case")
	// }

	// var ageGroup string = "g" //возрастные группы: "a", "b", "c"

	// switch ageGroup {
	// case "a", "b", "c":
	// 	fmt.Println("AgeGroup 10-40")
	// case "d", "e":
	// 	fmt.Println("AgeGroup 50-70")
	// default:
	// 	fmt.Println("You are too yang/old")
	// }

	// var age int
	// fmt.Scan(&age)

	// switch {
	// case age <= 18:
	// 	fmt.Println("Too yong")
	// case age > 18 && age <= 30:
	// 	fmt.Println("Second case")
	// case age > 30 && age <= 100:
	// 	fmt.Println("Too old")
	// default:
	// 	fmt.Println("Who are you")
	// }

	var number int
	fmt.Scan(&number)
	switch {
	case number < 100:
		fmt.Printf("%d is less then 100\n", number)
		fallthrough
	case number < 200:
		fmt.Printf("%d is less then 200\n", number)
	}

}
