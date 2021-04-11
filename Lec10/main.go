package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	name := "Hello world"
	fmt.Println(name)

	// строка - это байтовый слайс со своими особенностями при обращении к обертываемому массиву
	// len - считает кол-во байтов, а не символов
	// строки в Go хранятся как наборы UTF-8 символов, каждый символ может занимать больше чем 1 байт

	word := "Sample word"
	// word := "Тестовая строка"

	fmt.Printf("String: %s\n", word)
	//какие значения байт находятся в слайсе word?
	fmt.Printf("Bytes: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%x ", word[i]) //%x - формат представления 16-тиричночного представления
	}
	fmt.Println()

	// доступ к отдельным символам
	fmt.Printf("Characters: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c ", word[i]) //%c - формат представления символа
	}
	fmt.Println()

	// Руна (Rune) встроенный тип в Go (alias над int32), позволяющий хранить единый, неделимый UTF символ
	// в не зависимости от того сколько байт он занимает
	kirilicWord := "Тестовая строка"
	runeSlice := []rune(kirilicWord) // преобразование слайса байт к слайсу рун

	fmt.Printf("Runes: ")
	for i := 0; i < len(runeSlice); i++ {
		fmt.Printf("%c ", runeSlice[i])
	}
	fmt.Println()

	//итерирование по строке с использованием рун
	for id, runeVal := range kirilicWord { // id - это индекс байта, с которого начинается руна runeVal
		fmt.Printf("%c start at position %d\n", runeVal, id)
	}

	// создание строки из слайса байт
	myByteSlice := []byte{0x40, 0x41, 0x42, 0x43}
	myStr := string(myByteSlice)
	fmt.Println(myStr)

	myDecimalByteSlice := []byte{100, 101, 102, 103} // можно использовать десятичное представление байт
	myDecimalStr := string(myDecimalByteSlice)
	fmt.Println(myDecimalStr)

	// Создание строки из слайса Рун
	runeHexSlice := []rune{0x45, 0x46, 0x47, 0x48}
	myStrFromRune := string(runeHexSlice)
	fmt.Println(myStrFromRune)

	//Руны как литералы
	runLiteralSlice := []rune{'v', 'a', 's', 'y', 'a'} // ' - символ обозначения руны
	myStrFromRuneLit := string(runLiteralSlice)
	fmt.Printf("data: %s, type:%T\n", myStrFromRuneLit, myStrFromRuneLit)

	//Длина и ёмкость строки
	// длина len() - количество байт в слайсе
	fmt.Println("Length of Вася:", len("Вася"), "bytes")
	// длина RuneCountInString - количество рун (символов)
	fmt.Println("Length of Вася:", utf8.RuneCountInString("Вася"), "runes")
	// вычисление емкости строки не предусмотренно т.к. строка базовый тип

	// Сравнение строк (равенство - неравенство) и конкатенация
	word1, word2 := "Вася", "Петя"
	if word1 == word2 {
		fmt.Println("Братаны")
	} else {
		fmt.Println("Ну такое")
	}

	word3 := word1 + word2
	fmt.Println(word3)

	// построитель (шаблон) строк
	firstName := "Alex"
	secondName := "Johnson"
	fullName := fmt.Sprintf("%s ### %s", firstName, secondName)
	fmt.Println(fullName)

	// Строку изменить нельзя, только сформировать новую
	// как вариант преобразовать строку к слайсу рун и там возможно обращение по индексу и соотв. изменение
	fullNameSlice := []rune(fullName)
	fullNameSlice[3] = 'Я'
	fmt.Println(string(fullNameSlice))

	// Полезные методы работы со строками
	// import "strings"
}
