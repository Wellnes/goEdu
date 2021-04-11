package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var name string
	input := bufio.NewScanner(os.Stdin)

	if input.Scan() { // команда захвата потока ввода и сохранения в буфер (захват идет до символа окончания строки \n) true - захват удался
		name = input.Text() // команда возвращения элементов, помещенных в буфер (string)
	}

	fmt.Println(name)
	msVvoda := strings.Fields(name) // создает слайс из строки с разделителями - пробелами
	for _, val := range msVvoda {
		fmt.Println(val)
	}
	fmt.Printf("slice vvoda %q \n", msVvoda)

	fmt.Println("For loop started:")
	for {
		if input.Scan() {
			result := input.Text()
			// input.Bytes() - возвращает слайс байтов
			if result == "" {
				break
			}
			fmt.Println("Input string is:", result)
		}
	}

	// Преобразование строкового литера к чему-нибудь числовому
	numStr := "10"
	numInt, _ := strconv.Atoi(numStr) // Atoi - Anything to Int
	fmt.Printf("%d is %T\n", numInt, numInt)

	//numInt64, _ := strconv.ParseInt(numStr, 10, 64)
	//numFloat32, _ := strconv.ParseFloat(numStr, 32) // это 64-разрядное число будет без проблем приводится к 32-разрядному

}
