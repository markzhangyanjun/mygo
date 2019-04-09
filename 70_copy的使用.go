package main

import "fmt"

func main(){
	srcslice := []int{1,2}
	dstslice := []int{6,6,6,6,6}

	copy(dstslice,srcslice)

	fmt.Println("dst=",dstslice) //dst= [1 2 6 6 6]

}
