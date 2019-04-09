package main //必须有一个main
import "fmt"

func main(){
	var a int = 10
	//每个变量有两层含义：变量的内存，变量的地址

	fmt.Printf("a=%d\n",a)//变量的内存
	fmt.Printf("&a=%v\n",&a)//变量的地址

	//保存变量的地址，需要指针类型 *int 保存int地址 **int保存*int地址
	//定义声名：定义就是特殊的声名
	//定义一个变量p 类型为*p
	var p *int
	p = &a //指针变量指向谁就把谁的地址赋值给指针变量
	fmt.Printf("p=%v,&a=%v\n",p,&a)

	*p = 666 //*p操作的不是p的内存，是P指向的内存（就是a）
	fmt.Printf("*p=%v,a=%v\n",*p,a)
}
