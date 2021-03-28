package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Это моя вторая программа! Я рад, что учусь в Специалисте!!")
	//Декларирование переменных

	var age int
	fmt.Println("My age is:", age)
	age = 32
	fmt.Println("Age after assignment:", age)

	var height int = 183
	fmt.Println("Мой рост:", height)

	var uid = 1245
	fmt.Println("My uid:", uid)

	var firstVar, secondVar int = 20, 30
	fmt.Printf("F%d S%d \n", firstVar, secondVar)

	var (
		personName string = "Bob"
		personAge  int    = 42
		personUid  int
	)

	fmt.Printf("Name:%s\nAge%d\nUID:%d \n", personName, personAge, personUid)

	count := 10
	fmt.Println("Count:", count)
	count++
	fmt.Println("Count:", count)

	width, lenth := 36.4, 12.8
	fmt.Printf("Minimum is dannux: %.2f\n", math.Min(width, lenth))

	var word1, word2, word3, word4 string = "имя", "твое", "мне", "знакомо"
	fmt.Println(word4, word3, word2, word1)
	fmt.Println(word3, word1, word4, word2)
	fmt.Println(word2, word4, word1, word3)
}
