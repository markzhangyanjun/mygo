package main

import "fmt"

func main(){
	var p *int
	p = nil
	fmt.Println("p=",p)
	//*p = 666 //错误，因为p没有合法指向

	var a int
	p = &a
	*p = 788
	fmt.Printf("p=%v,a=%d",&a,*p)


}