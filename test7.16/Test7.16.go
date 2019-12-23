package main

import "fmt"

func main() {
	a :=[]int{21,54,96,3184,21,54,13,54,1,20,2}
	fmt.Println(sorts(a))
}
func sorts(a []int) ([]int){
	for i:=0;i<len(a);i++  {
		for j:=i+1;j<len(a);j++ {
			if a[i]<a[j] {
				temp:=a[i]
				a[i]=a[j]
				a[j]=temp
			}
		}
	}

	return  a
}
