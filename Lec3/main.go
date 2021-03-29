package main

import "fmt"

func main() {

	//	var (
	//		age  int
	//		name string
	//	)

	//fmt.Scan(&age)
	//fmt.Scan(&name)

	//fmt.Scan(&age, &name)
	//fmt.Printf("My name is: %s\nMy age is: %d\n", name, age)

	//fmt.Fscan(os.Stdin, &age)
	//fmt.Println("New age:", age)

	// var (
	// 	name, familia string
	// 	vozrast       int
	// )

	// fmt.Println("Enter name, familia, vozrast")
	// fmt.Scan(&name, &familia, &vozrast)
	// fmt.Printf("Имя: %s , Фамилия: %s , Возраст: %d . Студент BPS\n", name, familia, vozrast)

	//tea, chicken, fries, 13:00

	// var drink, meal, subMeal, time string

	// fmt.Scan(&drink, &meal, &subMeal, &time)
	// fmt.Printf("I wanna some '%s', my favorite meal - '%s'. Give me also '%s'. I will come soon... '%s'\n", drink, meal, subMeal, time)

	// var sl1, sl2, sl3, sl4, avatar string
	// fmt.Scan(&sl1, &sl2, &sl3, &sl4, &avatar)
	// sep := "-"

	// fmt.Println(sl4, sep, avatar)
	// fmt.Println(sl3, sep, avatar)
	// fmt.Println(sl2, sep, avatar)
	// fmt.Println(sl1, sep, avatar)

	// var aStorona, bStorona int
	// fmt.Scan(&aStorona, &bStorona)
	// fmt.Println("Периметр:", 2*(aStorona+bStorona))
	// fmt.Println("Площадь:", aStorona*bStorona)

	var chislo1, chislo2, sumChisel int
	fmt.Scan(&chislo1, &chislo2)
	sumChisel = chislo1 + chislo2
	fmt.Println(sumChisel * sumChisel)

}
