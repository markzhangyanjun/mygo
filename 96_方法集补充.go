package main

import "fmt"

type mint int
func (tmp mint) test_w(){

	fmt.Println("test_w:tmp=",tmp)

}

func(tmp *mint)test_m(){

	fmt.Println("test_m:tmp=",tmp)

}
func main(){

	var a mint = 10
	a.test_w()
	a.test_m()






}