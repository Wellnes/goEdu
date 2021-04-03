package main

import (
	"fmt"
)

func main() {

	// Задание С
	// var chVvod uint8 = 1

	// for chVvod != 0 {
	// 	fmt.Scan(&chVvod)
	// 	if chVvod != 0 {
	// 		fmt.Println(chVvod)
	// 	}
	// }

	// Задание D
	var strVvod string

mainfor:
	for {
		strVvod = ""
		fmt.Scanln(&strVvod)
		if len(strVvod) == 0 {
			break mainfor
		} else {
			fmt.Println(strVvod)
		}

	}

}
