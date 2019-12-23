package main

import "fmt"

func main() {
	s :=[]int{1,2,3,4,5,6,7,8,9,10}
	s1 :=Filter(s,fn)
	fmt.Printf("%d",s1)
}
func Filter(s []int,fn func(int) bool) []int{
	as :=make([]int,0,len(s))
	for _,a :=range s {
		if fn(a){
			as = append(as, a)
		}
	}
	return as
}

func fn( a int ) bool {
	if a%2==0 {
		return true
	}
	return false
}
