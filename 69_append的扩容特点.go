package main

import "fmt"

func main(){
	//如果超过原来的容量，通常以2倍容量扩容

	s := make([]int,0,1)

	oldcap := cap(s)

	for i := 0; i < 8; i ++{
		s = append(s,i)
		if newcap :=cap(s);oldcap < newcap{
			fmt.Printf("cap: %d====>%d\n",oldcap,newcap)
			oldcap = newcap
		}
	}

}