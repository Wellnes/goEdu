package main

import "fmt"

// Указатели - переменная хранящая в качестве значения адрес в памяти другой переменной
// zeroValue имеет значение nil (это указатель в никуда)
func main() {

	// определение указателя на что-то
	var variable int = 30
	var pointer *int = &variable // & ... - операция взятия адреса в памяти

	fmt.Println(pointer)
	// разыменование указателя ( получение значения )
	var numericValue int = 32
	var pointerToNumeric *int = &numericValue

	fmt.Printf("Value in numericValue is %v\nAddress is %v\n", *pointerToNumeric, pointerToNumeric)

	// создание указателей на нулевые значения типов
	zeroPoint := new(int) // создает zeroValue для int и возвращает адрес где это храниться
	*zeroPoint += 40

	// передача указателей в функции
	sample := 1
	samplePointer := &sample
	changeParam(samplePointer)
}

// определение функции принимающей параметр как указатель
func changeParam(val *int) {
	*val += 100
}
