package main

import "fmt"

func main() {
	a := 10
	b := 20

	defer func(){
		fmt.Printf("内部：a=%d,b=%d\n",a,b)
	}()
	a = 111
	b = 222
	fmt.Printf("外部：a=%d,b=%d\n",a,b)
}
