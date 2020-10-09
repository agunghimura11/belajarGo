package main

import "fmt"

type Sisi struct {
	Panjang, Lebar float32
}

func main() {
	first_name, _ := getFullName()
	fmt.Println(first_name)

	total := SumAll(10, 20, 30, 40)
	fmt.Println(total)

}

func getFullName() (FirstName, LastName string) {
	FirstName = "Ari"
	LastName = "Sudana"
	return
}

func SumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
