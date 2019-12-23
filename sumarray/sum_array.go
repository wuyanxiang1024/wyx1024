package main

import "fmt"

func main() {
	arrf :=[...]float32{1.0,2.0,3.0,4.0,5.5,6.5}
	//arr:=Sum(arrf)
	//fmt.Print(arr)
	slice :=arrf[0:]
	arr:=Sum(slice)
	fmt.Print(arr)
}
func Sum(arrf []float32) float32  {
	var arrs float32
	for _,arr := range arrf{
		arrs+=arr
	}
	return arrs
}
//func Sum(arrf [6]float32) float32  {
//	var arrs float32
//	for _,arr := range arrf{
//		arrs+=arr
//	}
//	return arrs
//}

