package main

import "fmt"

func chet(dataVal int) bool {

	rez := false

	if dataVal%2 == 0 {
		rez = true
	}

	return rez

}

func main() {

	//Классический условный оператор
	// if condition1 {  -- condition - должен возвращать булев тип
	// 	//body
	// } else if condition2 {
	//
	// } else {
	//
	//}

	// var value int
	// fmt.Scan(&value)

	// if chet(value) {
	// 	fmt.Println("Число", value, "четное")
	// } else {
	// 	fmt.Println("Число", value, "НЕ четное")
	// }

	// var color string
	// fmt.Scan(&color)
	// if strings.Compare(color, "зеленый") == 0 {
	// 	fmt.Println("Цвет зеленый")
	// } else if strings.Compare(color, "красный") == 0 {
	// 	fmt.Println("Цвет красный")
	// } else {
	// 	fmt.Println("Не ожидаемый цвет")
	// }

	// var ch1, ch2, ch3 string
	// fmt.Scan(&ch1, &ch2, &ch3)
	// fmt.Print(ch3, ch2, ch1, "\n")

	// var ch1, ch2 float32
	// fmt.Scan(&ch1, &ch2)

	// sumCh := int32(ch1 + ch2)
	// if sumCh%2 == 0 {
	// 	fmt.Println("ЧЁТНОЕ")
	// } else {
	// 	fmt.Println("НЕЧЁТНОЕ")
	// }

	// var login, email string
	// fmt.Scan(&login, &email)

	// const sobaka = "@"
	// const tochka = "."

	// if utf8.RuneCountInString(login) < 10 || strings.Contains(login, sobaka) {
	// 	fmt.Println("Некорректный логин")
	// } else if !(strings.Contains(email, sobaka) && strings.Contains(email, tochka)) {
	// 	fmt.Println("Некорректная почта")
	// } else {
	// 	fmt.Println("ОК")
	// }

	var x1, y1, x2, y2 uint8
	fmt.Scan(&x1, &y1, &x2, &y2)

	isX2GoodGorizont := (x2 == x1+2 || x2 == x1-2)
	isY2GoodGorizont := (y2 == y1+1 || y2 == y1-1)

	isX2GoodVertical := (x2 == x1+1 || x2 == x1-1)
	isY2GoodVertical := (y2 == y1+2 || y2 == y1-2)

	if (isX2GoodGorizont && isY2GoodGorizont) || (isX2GoodVertical && isY2GoodVertical) {
		fmt.Println("ДА")
	} else {
		fmt.Println("НЕТ")
	}

}
