package main

import (
	"fmt"
	"time"
)
//3.var ehere = log.Print(...)
func main() {
	start:=time.Now()
	f:=closre()
	for i:=0;i<6 ;i++  {
		fmt.Println(f())
	}
	end:=time.Now()
	delta :=end.Sub(start)
	println("执行时间",delta)
}

func closre() func()int{
/*	1.where :=func(){
		_,file,line ,_:=runtime.Caller(1)
		println(file,line)
	}
	2.where()
	log.SetFlags(1)
	log.Print("sssss")
*/
	x1,x2,sum :=0,1,0
	return func()int {
		sum=x1+x2
		x1=x2
		x2=sum
		return  sum
	}
}
