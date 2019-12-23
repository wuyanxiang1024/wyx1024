package main

import "fmt"

func main() {
	// slice是指向底层的数组
	var a []int //最简单的slice声明 无指定长度的声明为slice声明  （可理解为指向数组的指针）
	fmt.Println(a)
	a2 :=[10]int{1,2,3,4,5,6,7,8,9,10}//数组
	a3 :=a2//声明有指定长度的数组声明到创建一个slice变量
	fmt.Println(a3)
	a4 :=a2[5:10]//截取数组索引为5到索引为9的五个到声明的slice中
	fmt.Println(a4,"\n")

	//make声明slice
	s1:=make([]int ,3,5)  //param1为指定类型slice,param2指定 初始化元素的个数，param3指定slice的初始化容量大小
	fmt.Println(len(s1),cap(s1),s1,"\n")

	//slice与底层数组的关系
	k:=[]byte{'a','b','c','d','e','f','g'}
	k2:=k[2:5]
	fmt.Printf("%c",k2)
	k3 := k[3:5]
	fmt.Printf("\n%c",k3)

	//Reslice:Reslice时索引以被slice的切片为准，索引不可超过sliced容量或底层数组长度（cap），越界会导致引发错误
	sl1 :=k[2:7]
	fmt.Printf("\n\n%c",sl1)
	sl2 :=sl1[1:5]
	fmt.Printf("\n%c",sl2)
	sl3 :=sl1[3:5]//则会从sl2中的原始底层数组获取元素，因slice是指向底层数组，又因为容量为初始化10，则会从底层数组取超出slice范围的容量大小内获取,如超出容量则报错
	fmt.Printf("\n%c\n",sl3)


	//Append：可在slice尾部追加元素，可将一个slice追加到另一个slice上
	//		  如果长度为超过原初始化容量，则地址不变，返回原slice，否则重新分配数组地址拷贝原数据到新分配的slice地址上，返回新的slice
	 ce :=make([]int ,3,5)
	fmt.Printf("\n\n%p",ce)
	ce=append(ce,1,2 )
	fmt.Printf("\n%v %p",ce,ce)
	ce=append(ce,1,2,3 ,4,5,6)//元素大于出事容量，重新分配一个存储地址，slice为新的
	fmt.Printf("\n%v %p\n\n",ce,ce)

	//slice指向的是一个底层数组 示例：
	i :=[]int{1,2,3,4,5,6}
	st1 :=i[2:5]
	st2 :=i[1:3]
	fmt.Printf("\n%v %p与%v %p\n",st1,st1,st2,st2)
	st1[0]=9
	fmt.Printf("%v %p与%v %p\n",st1,st1,st2,st2)
	st2=append(st2,1,2,3 ,4,5,6)//当追加了元素，超出容量，重新分配的slice则指向新的地址，则共享的部分不在同步修改
	fmt.Printf("%v %p与%v %p\n",st1,st1,st2,st2)
	st2[1]=3//修改指向底层数组的值，则共享该值的其他，也跟着改变。如在修改前，改变了slice的地址的地址，则不在共享append时修改的同步
	fmt.Printf("%v %p与%v %p\n",st1,st1,st2,st2)
	println()

	//copy：将j2拷贝到j1
	j1 := []int{1,2,3,4,5,6}
	j2 := []int{7,8,9,10,11,12,13,14,15,16,17}
	fmt.Printf("%v %p与%v %p\n",j1,j1,j2,j2)
	copy(j1,j2)//将j2拷贝到j1
	fmt.Printf("%v %p与%v %p\n\n",j1,j1,j2,j2)
	//copy：将j2拷贝到j1
	j11 := []int{1,2,3,4,5,6}
	j22 := []int{7,8,9,10,11,12,13,14,15,16,17}
	fmt.Printf("%v %p与%v %p\n",j11,j11,j22,j22)
	copy(j22,j11)//将j1拷贝到j2
	fmt.Printf("%v %p与%v %p\n\n",j11,j11,j22,j22)

	//copy其中一部分指定范围到指定位置
	j111 := []int{1,2,3,4,5,6}
	j222 := []int{7,8,9,10,11,12,13,14,15,16,17}
	fmt.Printf("%v %p与%v %p\n",j111,j111,j222,j222)
	copy(j111[2:4],j222[2:4])//
	fmt.Printf("%v %p与%v %p\n\n",j111,j111,j222,j222)


	j3 :=append(j111)
	fmt.Printf("%v %p与%v %p\n",j111,j111,j3,j3)





}