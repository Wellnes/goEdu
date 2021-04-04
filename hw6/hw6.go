package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
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
	// 	var strVvod string

	// mainfor:
	// 	for {
	// 		strVvod = ""
	// 		fmt.Scanln(&strVvod)
	// 		if len(strVvod) == 0 {
	// 			break mainfor
	// 		} else {
	// 			fmt.Println(strVvod)
	// 		}

	// 	}

	//LOLKEK9000

	var (
		pass1, pass2             string
		shortPass                int    = 8
		simplePass1, simplePass2 string = "123", "qwe"
	)

passFor:
	for {

		fmt.Scan(&pass1, &pass2)

		if utf8.RuneCountInString(pass1) < shortPass {
			fmt.Println("Слишком короткий пароль!")
		} else if strings.Contains(pass1, simplePass1) || strings.Contains(pass1, simplePass2) {
			fmt.Println("Слишком простой пароль!")
		} else if pass1 != pass2 {
			fmt.Println("Введенные пароли различаются!")
		} else {
			fmt.Println("Ну наконец-то!")
			break passFor
		}

	}

}
