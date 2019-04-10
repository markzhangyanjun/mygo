package test

import "fmt"

//如果是首字母小写，只能在同一个包里使用
type  stu struct {
	id int
}

type  Stu struct {
	//id int  //如果是首字母小写，只能在同一个包里使用
	Id int
}



//如果是首字母小写，只能在同一个包里使用
func myFunc(){
	fmt.Println("this is myFunc")
}

//首字母答谢，外界也可以访问
func MyFunc(){
	fmt.Println("this is MyFunc")
}



