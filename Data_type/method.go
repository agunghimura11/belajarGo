package main

import "fmt"

type Sisi struct {
	Panjang int
	Lebar   int
}

func main() {
	s := Sisi{Panjang: 3, Lebar: 4}
	s.ScaleUp(2)
	fmt.Println(s.Luas())
}

func (s Sisi) Luas() int { // method
	return s.Panjang * s.Lebar
}

func (s *Sisi) ScaleUp(i int) { //pointer
	s.Lebar = s.Lebar * i
	s.Panjang = s.Panjang * i
}
