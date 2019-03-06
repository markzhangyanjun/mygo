package main

import "fmt"
//1.iota常量自动生成器，每一行自动累加1
//2.iota 给常量赋值使用
func main() {
	const (
		a= iota
		b= iota
		c= iota
		d= iota
	)

	fmt.Println("a=", a, "b=", b, "c=", c, "d=", d)
	//3.iota遇到const,重置为0
	const e= iota
	fmt.Println("e =", e)

	const (
		a1= iota
		b1
		c1
	)
	fmt.Println("a1=", a1, "b1=",b1,"c1=",c1)

	const(
		i = iota
		m,n,k = iota,iota,iota
		j = iota

	)
	fmt.Printf("i = %d ,m = %d ,n = %d ,k = %d, j = %d",i,m,n,k,j)
}
