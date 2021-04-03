package main

import "fmt"

func main() {

	// var b, a uint8 = 50, 1

	// for a < b {
	// 	a *= 2
	// 	fmt.Println("Now a is", a)
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Now i is", i)
	// }

	//Елочка
	// for i := 0; i < 10; i++ {
	// 	for j := 0; j <= i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Print("\n")
	// }

	// Лейблы
outer:
	for i := 0; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j=%d\n", i, j, i+j)
			if i == j {
				break outer // детерминирует внешний цикл
			}
		}
	}

}
