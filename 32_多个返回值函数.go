package main

import "fmt"

func fu1()(int,int,int){
	return 1,2,3
}
func fu2()(a int,b int,c int){
	a,b,c = 111,222,333
	return
}

func fu3()(a,b,c int){
	a,b,c = 444,555,666
	return
}

func main(){
	a,b,c :=fu1()
	fmt.Printf("a=%d\nb=%d\nc=%d\n",a,b,c)

	d,e,f := fu2()
	fmt.Printf("d=%d\ne=%d\nf=%d\n",d,e,f)

	g,h,i := fu3()
	fmt.Printf("g=%d\nh=%d\ni=%d\n",g,h,i)


}
