package main

import (
	"ControlSystem"
	"calc"
	"fmt"
)

func hello() {
	fmt.Println("TESTTEST")
	fmt.Println(calc.Sum(10, 20))
	fmt.Println((ControlSystem.Test(10)))
}

func main() {
	go hello()

	fmt.Scanln()
}
