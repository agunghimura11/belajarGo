package main

import (
	"fmt"
	"math"
)

func main() {
	// arr := []string{
	// 	"satu",
	// 	"dua",
	// 	"tiga",
	// 	"empat",
	// }
	// perulangan(arr)

	fmt.Println(sqrt(2), sqrt(-4))
}

func perulangan(arr []string) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i, v := range arr {
		fmt.Println(fmt.Sprintf("Value %s %y", i, v))
	}

}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
