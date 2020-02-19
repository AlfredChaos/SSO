package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type test struct {
}

func print() int {
	fmt.Println("test ok")
	return 1
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
	pv := register[1].Call(nil)
	fmt.Println(pv[0])
}

func prints(i int) string {
	fmt.Println("i=", i)
	return strconv.Itoa(i)
}
