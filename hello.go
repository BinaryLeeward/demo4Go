package main

import "fmt"
import "reflect"

type X struct {
	i int
	j string
}

func main() {
	fmt.Println("hello binaryleeward!")
	a := new(X)
	b := X{1, "bbb"}
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Printf("a: %v %T \n", a, a)
	fmt.Printf("b: %v %T \n", b, &b)
	fmt.Println(&a)
	fmt.Println(&b)
}
