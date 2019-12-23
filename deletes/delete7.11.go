package main

import "fmt"

func main() {
	//a:=[]int{1,2,3,4,5,6}
	//a2 := make([]int,6,20)
	//copy(a2[2:6],a)
	//a =append(a[:0])
	//fmt.Printf("%d\n",a2)
	//fmt.Printf("%d",a)
	s:="\u00ff\u754c"
	var ss []byte
	for _, c := range s {
		fmt.Printf("%c",c)
	}
	fmt.Println("\n",copy(ss,s))

}
