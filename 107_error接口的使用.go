package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 :=fmt.Errorf("this is normal err")
	fmt.Println("err1=",err1)

	err2 :=errors.New("this is normal err2")
	fmt.Println("err2=",err2)
}
