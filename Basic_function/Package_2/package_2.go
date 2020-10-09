package Package_2

import "fmt"

func PrintFromPackage1() {
	fmt.Println("Ini adalah package 1")
}

func CustomPrint(str string) {
	fmt.Println(fmt.Sprintf("Hallo nama saya $s", str))
}
