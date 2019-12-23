package main

import (
	"fmt"
)

func main() {
	sl :=make([]byte,2,2)
	s2 :=make([]byte,6,6)
	fmt.Println(sl,s2)
	s3 :=Append(sl,s2)
	fmt.Println(len(s3),cap(s3))
	fmt.Println(s3)
	//items :=[...]int{10,20,30,40}
	//for _,item:=range items {
	//	item*=2
	//	fmt.Println(item)
	//}
}
func Append(slice ,data []byte )[]byte {
	l:=len(slice)
	if l+len(data)>cap(slice){
		newSlice := make([]byte,(l+len(data))*2)
		for s,i :=range slice{
			newSlice[s]=i
		}
		slice=newSlice[:(l+len(data))]
		for s,i:= range data{
			slice[s+l]=i
		}
	}
	slice =slice[0:l+len(data)]
	for i,s :=range data{
		slice[l+i]=s
	}

	return slice
}
