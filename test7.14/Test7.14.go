package main

import "fmt"

var str="wuyanxiang"
func main() {
	fmt.Println(Reverse(str))
}
func Reverse(str string )string{
	slice :=[]byte(str)
	for i,j :=0,len(slice)-1;i<len(slice)/2;i,j =i+1,j-1{
		slice[i],slice[j]=slice[j],slice[i]
	}
	return string(slice)

}