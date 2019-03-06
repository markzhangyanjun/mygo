package main

import "fmt"

func main(){
	var a int
	var b float64
	a , b = 10, 3.14
	fmt.Println("a=",a)
	fmt.Println("b=",b)

	var(
		c int
		d float64
	)
	c ,d = 10 ,3.14
	fmt.Printf("c = %d,d = %d\n",c,d)

	const i int = 10
	const j  float64 = 3.14
	fmt.Println("i=",i,"j=",j)

	const(
		m int = 100
		n float64 = 3.14
	)
	fmt.Println("m=",m,"n=",n)
}
