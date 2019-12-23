package main

import "fmt"

var a =make([]int,0)
var factor int=5

func main() {
	fmt.Println(len(a))
	a = lendou(a,factor)
	fmt.Println(len(a))
}
func lendou(a []int,b int) []int {
	if len(a)==0 {
		a = make([]int,1)
		return a
	}
	a =make([]int,len(a)*b)
	return a
}