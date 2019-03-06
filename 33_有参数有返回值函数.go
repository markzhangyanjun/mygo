package main

import "fmt"

func MaxandMin(a,b int)(max,min int){
	if a > b{
		max = a
		min = b
	}else{
		max = b
		min = a
	}
	return
}


func main(){
   max,min := MaxandMin(10,20)
   fmt.Printf("max=%d\nmin=%d",max,min)
}
