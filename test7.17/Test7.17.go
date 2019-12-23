package main

import "fmt"

func main() {
	a :=[]int{1,2,3,4,5,6}
	a2:=mapfun(parames,a)
	fmt.Println(a2)
}

func mapfun(f func(int)int ,list []int)[]int  {
	res :=make([]int,len(list))
	for i,v :=range list{
		res[i]=f(v)
	}
	return res
}
func parames(x int)int  {
	return x*10
}
