package main

import "fmt"

func main() {

	var KvoVvodimux, iChislo, startIndex, endIndex, TotalSum int

	fmt.Scan(&KvoVvodimux)
	slChisel := []int{}

	for i := 0; i < KvoVvodimux; i++ {
		fmt.Scan(&iChislo)
		slChisel = append(slChisel, iChislo)
	}

	fmt.Scan(&startIndex)
	startIndex--

	fmt.Scan(&endIndex)
	endIndex--

	for id, val := range slChisel {
		if startIndex <= id && id <= endIndex {
			TotalSum += val
		}
	}

	fmt.Println(TotalSum)

}
