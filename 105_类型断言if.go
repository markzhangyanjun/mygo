package main

import "fmt"

type Student_c struct {
	name string
	id  int
}


func main() {
	i :=make([]interface{}, 3)
	i[0] = 1
	i[1] = "hello"
	i[2] = Student_c{"mike",666}

	//类型查询，类型断言
    //第一个返回下标，第二个返回下标对应的值
	for index , data :=range i{

		//fmt.Println("x=",index,"value=",data)
		//第一个返回值：接口变量本身，第二个返回判断结果的真假
		if value ,ok :=data.(int);ok==true{

			fmt.Printf("x[%d]类型为int,内容为%d\n",index,value)

		} else if value ,ok :=data.(string);ok==true{

			fmt.Printf("x[%d]类型为string,内容为%d\n",index,value)

		} else if value ,ok :=data.(Student_c);ok==true{

			fmt.Printf("x[%d]类型为Student_c,内容为 %v\n",index,value)
		}

	}

}
