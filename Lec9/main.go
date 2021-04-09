package main

import "fmt"

func main() {

	// Слайсы (они же - срезы)
	// Слайс - это динамическая обвязка над массивом

	startArr := [4]int{10, 20, 30, 40}
	// Слайс определяется пустыми квадратными собками
	// создаем слайс на основании существующего массива
	var startSlice []int = startArr[0:2] // элементы с индексом от 0 до 2 не включительно [:] - все элементы
	fmt.Println("Slice[0:2]:", startSlice)

	// Создание слайса без явной инициализации массива
	secondSlice := []int{15, 20, 30, 40}
	fmt.Println("SecondSlice:", secondSlice)

	// Изменение элементов среза
	originArr := [...]int{30, 40, 50, 60, 70, 80}
	firstSlice := originArr[1:4] // набор ссылок на элементы нижележащего массива
	for i, _ := range firstSlice {
		firstSlice[i]++
	}
	fmt.Println("OriginArr:", originArr)
	fmt.Println("firstSlice:", firstSlice)

	// Один массив и два производных среза
	fSlice := originArr[:]
	sSlice := originArr[2:5]

	fmt.Println("Before modifications: Arr:", originArr, "fSlice", fSlice, "sSlice", sSlice)
	fSlice[3]++
	sSlice[1]++
	fmt.Println("After modifications: Arr:", originArr, "fSlice", fSlice, "sSlice", sSlice)

	// Срез как встроенный тип
	// type slice struct {
	// 	Length int
	// 	???? int
	// 	ZeroElement *byte
	// }

	// Длина и емкость слайса
	wordsSlice := []string{"one", "two", "three"}
	fmt.Println("slice:", wordsSlice, "Length:", len(wordsSlice), "Capasity", cap(wordsSlice))
	// Capasity - размер охватываемого слайсом массива
	// возвращает кол-во элементов, которые можно хранить в этом слайсе без выделения доп. памяти для новых элементов
	// охватываемый массив увеличивается сразу на n*2 где n - первоначальный объем массива
	// как только увеличивается Capasity ссылки с первым массивом разрываются
	wordsSlice = append(wordsSlice, "four")
	fmt.Println("slice:", wordsSlice, "Length:", len(wordsSlice), "Capasity", cap(wordsSlice))

	// make() функция позволяющая создавать в том числе и срезы
	// []int - тип коллекции
	// 10 - длина
	// 15 - Capasity
	sl := make([]int, 10, 15)

	// добавление одного слайса в конец другого
	// можно через цикл, но есть
	anotherSlice := []int{7, 8}
	sl = append(sl, anotherSlice...)
}
