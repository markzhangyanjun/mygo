package main

import "fmt"
//切片不是数组
func main(){
	a := []int{1,2,3,0,0,2,4,9,8}
	s :=a[0:3:8]

	fmt.Println("s=",s)

	fmt.Println("len(s)=",len(s)) //3-0
	fmt.Println("cap(s)=",cap(s))//8-0
}
