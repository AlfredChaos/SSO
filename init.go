package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type test struct {
}

func print(i int) {
	fmt.Println("test ok")
	fmt.Print(i)
	//return 1
}

func main() {
	//data.UserList = append(data.UserList, data.User1)
	//t := &test{}
	register := make(map[int]reflect.Value)
	register[1] = reflect.ValueOf(print)
	/*for k, v := range register {
		fmt.Println(k)
		fmt.Println(v)
	}*/
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(1)
	register[1].Call(params)
	//fmt.Println(pv[0])
}

func prints(i int) string {
	fmt.Println("i=", i)
	return strconv.Itoa(i)
}
