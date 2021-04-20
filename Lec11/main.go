package main

import (
	"fmt"
)

func main() {

	// map - предназначен для хранения пар ключ:значение (аналог соответствия в 1С)
	// инициализация
	mapper := make(map[string]int)
	fmt.Println("Empty map:", mapper)

	//добавление пар в существующее соответсвие
	mapper["Alice"] = 24
	mapper["Bob"] = 25
	fmt.Println("Map after adding pairs:", mapper)

	// инициализация с указанием пар
	newMapper := map[string]int{
		"Alice": 1000,
		"Bob":   1000,
	}
	newMapper["Jo"] = 1000
	fmt.Println("New mapper ", newMapper)

	// что может быть ключем?
	// 1. ВАЖНО: НЕ упорядочена в Go
	// 2. ключем должен быть только сравнимый между собой тип (==, !=)

	// получение элементов
	testPerson := "Alice"
	fmt.Println("Salary of testPerson:", newMapper[testPerson])

	// получение элементов, которых нет в соответсвии - вернет значение по умолчанию для типа значения
	testPerson = "Derek"
	fmt.Println("Salary of testPerson:", newMapper[testPerson]) // при обращении по несуществующему ключу новая пара не добавляется

	// проверка вхождения ключа: при обращении по ключу возвращаются 2 значения
	employee := map[string]int{
		"Den":   100,
		"Alice": 200,
		"Bob":   300,
	}

	if value, ok := employee["Den"]; ok {
		fmt.Println("Den and value:", value)
	} else {
		fmt.Println("Den does not exists in map")
	}

	// перебор элементов
	for key, value := range employee {
		fmt.Printf("%s and value %d\n", key, value)
	}

	// удаление элементов
	delete(employee, "Den")

	// кол-во пар - это длина map
	fmt.Println("Pair amount in map:", len(employee))

	// map (как и slice) это ссылочный тип

}
