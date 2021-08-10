package main

import "fmt"

var arr = [3]int{6, 6, 6}

func IncFc(x int) int {
	return x + 1
}

func main() {
	fmt.Printf("Hello YuLong! \n")
	fmt.Println(arr)
	// Variable as Functions
	var FuncVar func(int) int // decalre FunVar as an function type variable : general declaration "var name func(type) type"
	FuncVar = IncFc
	fmt.Println(FuncVar(1))
}
