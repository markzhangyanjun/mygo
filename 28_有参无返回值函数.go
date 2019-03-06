package main

import "fmt"

func MyFunc01(a int){
	//a = 111
	fmt.Println("a=",a)
}

func Myfunc02(a int,b int){
	fmt.Printf("a=%d b=%d",a,b)
}

func Myfunc03(a,b int){
	fmt.Printf("a=%d b=%d",a,b)
}

func Myfunc04(a int,b string,c float64){
	fmt.Printf("a=%d , b=%s ,c=%f",a,b,c)
}

func Myfunc05(a,b string,c,d int,e,f float32){
	fmt.Printf("a=%s b=%s\nc=%d d=%d\ne=%f f=%f\n",a,b,c,d,e,f)
}

func main(){
	MyFunc01(12)
	Myfunc05("a","b",1,2,3.14,3.14)
}
