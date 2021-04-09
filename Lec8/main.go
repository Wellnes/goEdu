package main

func main() {

	// Массивы
	// статичен, хранит объекты одного типа

	//Создадим массив для хранения 5-ти целочисленных элементов
	// var arr [5]int
	// fmt.Println("This is my array:", arr)

	// arr[0] = 10
	// arr[1] = 20
	// arr[3] = -500
	// arr[4] = 720

	// fmt.Println("After elements init:", arr)

	//Определение массива с инициализацией
	// newArr := [5]int{10, 20, 30, 40, 50}
	// fmt.Println("Short declaration and init", newArr)

	// arrVal := [...]int{11, 21, 31, 41}
	// fmt.Println("Fashion Declar", arrVal, "Length:", len(arrVal))

	//Итерирование по массиву
	// floatArr := [...]float64{12.5, 13.5, 15.2, 10.0, 12.0}
	// for i := 0; i < len(floatArr); i++ {
	// 	fmt.Printf("%d element of arr is %.2f\n", i, floatArr[i])
	// }

	//Итерирование по массиву через оператор range
	//range возвращает 2 значения: текущий индекс элемента и его значение
	// floatArr := [...]float64{12.5, 13.5, 15.2, 10.0, 12.0}
	// for id, val := range floatArr {
	// 	fmt.Printf("%d element of arr is %.2f\n", id, val)
	// }

	//Игнорирование id в range
	// for _, val := range floatArr {
	// 	fmt.Printf("element of arr is %.2f\n", val)
	// }

	//Многомерные массивы
	// words := [2][2]string{
	// 	{"Bob", "Alice"},
	// 	{"Victor", "Jo"},
	// }

	// fmt.Println("Multidimensional array:", words)

	// for _, curMassa := range words {
	// 	for _, val := range curMassa {
	// 		fmt.Printf("%s ", val)
	// 	}
	// 	fmt.Println()
	// }

	// Срез
	// slice := []int{10,20,30}
	// slice[0] = 10
	// slice = append(slice, 200) // добавление нового элемента

	// var serServiceCount, serMyListCount int
	// fmt.Scan(&serServiceCount)
	// fmt.Scan(&serMyListCount)

	// sliceSerService := []string{}
	// var filmName string = "ara"

	// for i := 0; i < 4; i++ {
	// 	fmt.Fscan(os.Stdin, &filmName)
	// 	fmt.Println(filmName)
	// }

	// for i := 0; i < serMyListCount; i++ {
	// 	sliceSerService = append(sliceSerService, filmName)
	// }

	// for i, v := range sliceSerService {
	// 	fmt.Println(i, v)
	// }

}
