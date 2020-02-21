package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
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
	/*register := make(map[int]reflect.Value)
	register[1] = reflect.ValueOf(print)*/
	/*for k, v := range register {
		fmt.Println(k)
		fmt.Println(v)
	}*/
	/*params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(1)
	register[1].Call(params)*/
	//fmt.Println(pv[0])
	/*var test = make(map[int]string, 1)
	test[0] = "test"
	if _, exist := test[0]; exist {
		fmt.Println("exist")
	}*/
	w := md5.New()
	_, _ = io.WriteString(w, "123456")
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	fmt.Println(md5str)
	fmt.Println(time.Now())
}

func prints(i int) string {
	fmt.Println("i=", i)
	return strconv.Itoa(i)
}
