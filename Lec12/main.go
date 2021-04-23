package main

import "fmt"

// функции
// явная функция - локально-определенный блок кода, имеющий имя
//
//func functionName( params type ) typeReturnValue {
//	//body
//}

func add(a int, b int) int {
	result := a + b
	return result
}

// функция с однотипными параметрами
func mult(a, b, c, d int) int {
	result := a * b * c * d
	return result
}

// возврат более одного значения
func rectungleParameters(length, width float64) (float64, float64) {

	var perimeter = 2 * (length + width)
	var area = length * width

	return perimeter, area
}

// именованный возврат значений
func namedReturn(a, b int) (perimeter, area int) {
	perimeter = 2 * (a + b) // := - вызовет ошибку т.к. уже определили выше
	area = a * b
	return // не нужно указывать возвращаемые значения
}

// Variadic parameters (континуальные параметры)
// при смешении параметров Variadic всегда последний
// Variadic на всю функцию один
// sliceNumber... -- распаковывает слайс в числа через запятую
func helloVariadic(a ...int) {
	fmt.Println(a)
}

func main() {

	rez := add(10, 20)
	fmt.Println("rez add:", rez)
	fmt.Println("result mult:", mult(1, 2, 3, 4))

	per, area := rectungleParameters(10.5, 10.5)
	fmt.Println("rez area:", area)
	fmt.Println("rez per:", per)

	secondPer, secondArea := namedReturn(10, 20)
	fmt.Println("rez secondPer:", secondPer)
	fmt.Println("rez secondArea:", secondArea)

	helloVariadic(10, 20, 30, 40, 50)

	// анонимная функция
	anon := func(a, b int) int {
		return a + b
	}
	fmt.Println("Anon:", anon(20, 30))

}
