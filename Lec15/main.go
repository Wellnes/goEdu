package main

import "fmt"

// Структура - заименованный набор полей(свойств, состояний), определяющий новый тип данных

// явное определение структуры
type Student struct {
	firstName string
	lastName  string
	age       int
}

type AnotherStudent struct {
	firstname, lastName, groupName string
	age, curseNumber               int
}

func PrintStudent(std Student) {
	fmt.Println("===================")
	fmt.Println("FirstName:", std.firstName)
	fmt.Println("LastName:", std.lastName)
	fmt.Println("Age:", std.age)
}

func main() {
	// создание представителей структуры
	stud1 := Student{
		firstName: "Fedya",
		lastName:  "Petrov",
		age:       21,
	}
	//fmt.Println("Student 1:", stud1)
	PrintStudent(stud1)

	stud2 := Student{"Petya", "Ivanov", 19} // порядок указания свойств такой же как в структуре
	PrintStudent(stud2)

	// анонимные структуры (структуры без имени)
	anonStudent := struct {
		ara     bool
		age, id int
	}{
		ara: true,
		age: 22,
		id:  1,
	}
	fmt.Println(anonStudent)

	// доступ к состояниям (атрибутам) осуществляется через точку
	fmt.Println("FirstName:", stud2.firstName)
	stud2.age += 2

	// инициализация пустой структуры
	emptyStudent := Student{}
	PrintStudent(emptyStudent)

	// указатели на экземпляры структур
	studPointer := &Student{
		firstName: "Igor",
		lastName:  "Sidorov",
		age:       22,
	}
	PrintStudent(*studPointer)
}
