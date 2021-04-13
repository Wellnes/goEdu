package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// КекФликс
func main() {

	scSerialuOnService := []string{}
	scSerialuOnMyMind := []string{}

	var onServiceCount, onMyMindCount int

	input := bufio.NewScanner(os.Stdin)

	slugebnuVvod := 2
	var iInput, fullInputCount, serviceInputCount int

	for {
		if input.Scan() {

			iInput++
			fullInputCount = slugebnuVvod + onMyMindCount + onServiceCount
			serviceInputCount = slugebnuVvod + onServiceCount

			scanResult := input.Text()
			if scanResult == "" || fullInputCount < iInput {
				break
			}

			if iInput == 1 {
				scanIntResult, _ := strconv.Atoi(scanResult)
				onServiceCount = scanIntResult
			} else if iInput == 2 {
				scanIntResult, _ := strconv.Atoi(scanResult)
				onMyMindCount = scanIntResult
			} else if iInput <= serviceInputCount {
				scSerialuOnService = append(scSerialuOnService, scanResult)
				// fmt.Println("Add on service: ", scanResult)
			} else if serviceInputCount < iInput && iInput <= fullInputCount {
				scSerialuOnMyMind = append(scSerialuOnMyMind, scanResult)
				// fmt.Println("Add on my mind: ", scanResult)
			}

			if fullInputCount == iInput {
				break
			}

		}
	} // Input loop

	var found bool
	count := len(scSerialuOnService)

	for _, valMind := range scSerialuOnMyMind {

		found = false

		for i := 0; (!found) && (i < count); i++ {
			found = valMind == scSerialuOnService[i]
			if found {
				break
			}
		}

		if found {
			fmt.Println("ЕСТЬ")
		} else {
			fmt.Println("НЕТ В НАЛИЧИИ")
		}

	} // Output loop

}
