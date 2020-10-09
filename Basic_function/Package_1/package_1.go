package Package_1

import "fmt"

func PrintFromPackage1() {
	fmt.Println("Ini adalah package 1")
}

func CustomPrint(str string) {
	fmt.Println(fmt.Sprintf("Hallo nama saya %s", str))
}

func add(x int, y int) int {
	return x + y
}

func Add(x int, y int) int {
	return x + y
}

func AddDoubleReturn(x int, y int) (int, bool) {
	return x + y, true
}
