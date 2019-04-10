package main

import "fmt"

type Student_d struct {
	name string
	id  int
}
func main() {

	i :=make([]interface{}, 3)
	i[0] = 1
	i[1] = "hello"
	i[2] = Student_d{"mike",666}

	//类型查询，类型断言
	for index, data :=range i{
		switch value :=data.(type)  {
		case int:
			fmt.Printf("x[%d]类型为int,内容为%d\n",index,value)

		case string:
			fmt.Printf("x[%d]类型为string,内容为%d\n",index,value)
		case Student_d:
			fmt.Printf("x[%d]类型为Student_c,内容为 %v\n",index,value)
		default:
			fmt.Println("都不对")
		}




	}

}