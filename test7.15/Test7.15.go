package main

import "fmt"

var str []byte=[]byte{'a','b','a','a','a','c','d','e','f','g'}
func main() {
	fmt.Printf("%c",printsss(str))
}
func printsss(str []byte) []byte{
	arr:=make([]byte ,len(str))
	i:=0
	tmp:=byte(0)
	for _,strs :=range str {
		if strs != tmp {
			arr[i]=strs
			i++
		}
		tmp=strs
	}
	return arr
}