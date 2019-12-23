package main

import "fmt"

var str ="hello world!"
func main() {
	str1,str2 :=sqlit(str,5)
	slice := str[len(str)/2:]+str[:len(str)/2]

	fmt.Println(str1,str2,"\n",slice)
}
func sqlit(str string ,index int )(string,string){
	c :=[]byte(str)
	slice :=make([]byte,0,len(c))
	slice = append(slice,c[:index]...)
	slice2 :=make([]byte,0,len(c))
	slice2 = append(slice2,c[index:]...)
	str1 :=string(slice)
	str2 :=string(slice2)
	return str1,str2
}
