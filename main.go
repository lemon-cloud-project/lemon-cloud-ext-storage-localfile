package main

import "fmt"

var num = 0

func Plus() int {
	num++
	fmt.Println("[plus]current now is :", num)
	return num
}

func Sub() int {
	num--
	fmt.Println("[sub]current now is :", num)
	return num
}
