package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Test struct {
	X int
	Y string
}

func main() {
	fmt.Println("hello world!")
	test1()
	test2()
}

func test1() {
	a := Test{}
	fmt.Printf("a: %v %T \n", a, a)
	fmt.Println(a)
	err := json.Unmarshal([]byte(`{"X":1,"Y":"x"}`), &a)
	checkError(err)
	fmt.Printf("a: %v %T \n", a, a)
}

func test2() {
	fmt.Println("===========================")
	m := make(map[string]reflect.Type)
	m["test"] = reflect.TypeOf(Test{})
	a := reflect.New(m["test"]).Elem().Interface()
	fmt.Printf("a: %v %T \n", a, a)
	fmt.Println(a)
	err := json.Unmarshal([]byte(`{"X":1,"Y":"x"}`), &a)
	checkError(err)
	fmt.Printf("a: %v %T \n", a, a)
}
