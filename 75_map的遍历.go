package main

import "fmt"

func main(){
	m := map[int]string{1:"mike",2:"zyj",3:"www"}

	//第一个返回值为key,第二个返回值为value，遍历是无序的
	for key,value :=range m{

		fmt.Printf("key=%d,value=%s\n",key,value)

	}

	//如何判断一个key值是否存在
	//第一个返回值为key所对应的value，第二个返回值为key是否存在的条件，存在ok为 True
	value ,ok :=m[1]
	fmt.Println(value)
	fmt.Println(ok)


	if ok == true{
		fmt.Println("m[1]=",value)
	}else{
		fmt.Println("key不存在")
	}



}
