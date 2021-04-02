package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

// func chet(dataVal int) bool {
//
// 	rez := false
//
// 	if dataVal%2 == 0 {
// 		rez = true
// 	}
//
// 	return rez
//
// }

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

	// var x1, y1, x2, y2 uint8
	// fmt.Scan(&x1, &y1, &x2, &y2)

	// isX2GoodGorizont := (x2 == x1+2 || x2 == x1-2)
	// isY2GoodGorizont := (y2 == y1+1 || y2 == y1-1)

	// isX2GoodVertical := (x2 == x1+1 || x2 == x1-1)
	// isY2GoodVertical := (y2 == y1+2 || y2 == y1-2)

	// if (isX2GoodGorizont && isY2GoodGorizont) || (isX2GoodVertical && isY2GoodVertical) {
	// 	fmt.Println("ДА")
	// } else {
	// 	fmt.Println("НЕТ")
	// }

	// Инициализация в блоке условного оператора
	// блок присваивания - только :=
	// инициализируемая переменная видна только в области видимости условного оператора
	// if data, err := SOME_FUNC(); err != nil {
	//	}
	// if num := 10; num%2 == 0 {
	// 	fmt.Println("Четное")
	// } else {
	// 	fmt.Println("Не четное")
	// }

	// var userMassage string
	// fmt.Scan(&userMassage)

	// BOT := "чат"

	// if strings.Contains(userMassage, BOT) {
	// 	fmt.Println("БОТ")
	// } else {
	// 	fmt.Println("НЕ БОТ")
	// }

	// var raz, dva, tri string

	// fmt.Scan(&raz, &dva, &tri)

	// OK := "ОК"

	// if (raz == "раз" || raz == "один") && dva == "два" && tri == "три" {
	// 	fmt.Println(OK)
	// } else if raz == "1" && dva == "2" && tri == "3" {
	// 	fmt.Println(OK)
	// } else {
	// 	fmt.Println("НЕ ПРАВИЛЬНО")
	// }

	var textLogo string
	fmt.Scan(&textLogo)

	var cenaSymbol float64 = 0.23
	sumLogo := float64(utf8.RuneCountInString(textLogo)) * cenaSymbol
	rubPart := math.Floor(sumLogo)
	copPart := (sumLogo - rubPart) * 100
	if rubPart == 0 {
		fmt.Printf("%.0f коп.\n", copPart)
	} else {
		fmt.Printf("%.0f р. %.0f коп.\n", rubPart, copPart)
	}

}
