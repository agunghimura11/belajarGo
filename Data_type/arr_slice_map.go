package main

import "fmt"

func main() {
	mapExample()
}

func array() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slices() {
	name := []string{
		"nama_1",
		"nama_2",
		"nama_3",
	}

	a := name[0:2]

	a[0] = "--"
	fmt.Println(name)
}

type Geo struct {
	Lat  float32
	Long float32
}

func mapExample() {
	mapGeo := make(map[string]interface{}) //interface untuk tipe bebas

	mapGeo["Bali"] = Geo{
		Lat:  7.7777,
		Long: 1100,
	}

	mapGeo["Jakarta"] = Geo{
		Lat:  6666,
		Long: 100.000,
	}

	for i, v := range mapGeo {
		fmt.Println(fmt.Sprintf("Kordinat $s : $v", i, v))
	}
}
