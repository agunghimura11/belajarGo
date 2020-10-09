package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{X: 3, Y: 5}

	p := &v

	p.X = 10

	//fmt.Println(v)

	a := &Vertex{X: 20, Y: 30}
	q := a
	q.X = 5

	//fmt.Println(*a)

	var arrStruck []Vertex
	for i := 0; i < 5; i++ {
		arrStruck = append(arrStruck, Vertex{X: i, Y: i + 1})
	}

	fmt.Println(arrStruck)

}
