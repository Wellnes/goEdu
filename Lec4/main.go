package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//Boolean => default false
	var firstBoolean bool
	fmt.Println(firstBoolean)
	//Boolean operands
	aBoolean, bBoolean := true, false
	fmt.Println("AND:", aBoolean && bBoolean)
	fmt.Println("OR:", aBoolean || bBoolean)
	fmt.Println("NOT:", !aBoolean)

	//Numerics. Integers
	// uint8       the set of all unsigned  8-bit integers (0 to 255)
	// uint16      the set of all unsigned 16-bit integers (0 to 65535)
	// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	// int8        the set of all signed  8-bit integers (-128 to 127)
	// int16       the set of all signed 16-bit integers (-32768 to 32767)
	// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	// float32     the set of all IEEE-754 32-bit floating-point numbers
	// float64     the set of all IEEE-754 64-bit floating-point numbers

	// complex64   the set of all complex numbers with float32 real and imaginary parts
	// complex128  the set of all complex numbers with float64 real and imaginary parts

	// byte        alias for uint8
	// rune        alias for int32
	var aInt = 32
	// Вывод типа через %T форматирование
	fmt.Printf("Type is %T\n", aInt)
	// Узнаем сколько байт занимает переменная типа int
	fmt.Printf("Type %T size of %d bytes\n", aInt, unsafe.Sizeof(aInt))

	//Явное привидение типов
	var first32 int32 = 12
	var second64 int64 = 13
	fmt.Println(first32 + int32(second64))

	//Строки, Строка - это набор байт
	firstName := "Федя"
	lastName := "Pupkin"
	fullName := firstName + " " + lastName
	fmt.Println(fullName)
	fmt.Println("Length of string:", len(firstName)) //string length in bytes
	fmt.Println("Amount of chars:", utf8.RuneCountInString(firstName))

	//Поиск подстроки в строке
	totalString, subString := "ABCDEDFG", "CDE"
	fmt.Println(strings.Contains(totalString, subString))

	var sampleRune rune = 'Q' // одинарные кавычки
	fmt.Printf("Rune as char %c \n", sampleRune)

}
