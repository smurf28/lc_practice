package main

import (
	"fmt"
)

func begin(funName string) string {
	fmt.Println("Enter Function")
	return funName
}
func end(funName string) string {
	fmt.Println("Leave Function")
	return funName
}
func record() {
	fmt.Println("Run Function")
}
func main() {
	defer end(begin("func"))
	record()
}
