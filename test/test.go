package main

import "fmt"

func main() {
	for i:=0;i<10 ;i++  {
		fmt.Println(funcss(i))
	}
}
func funcss(s int ) int  {
	if s<=1 {
		s = 1
	}else {
		s = funcss(s-1)+funcss(s-2)
	}
	return s
}
