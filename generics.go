package main

import "fmt"

func add[T int | float32](i,j T) T {
	return i+j
}

type test struct{
	i int
}

func alter(x *[]int) {
	(*x)[0] = 3
}

func main() {
	fmt.Println(add[int](1,2))
	fmt.Println(add[float32](2.5, 3.2))
	
	x := make([]int, 0, 2)
	x = append(x,1)
	
	y := append([]int{},x...)
	
	for i:=0;i<len(x);i++{
		x[i]=0
	}
	
	fmt.Println(x,y)
	
	var t *test = &test{1}
	fmt.Println(*t)
	
	alter(&x)
	fmt.Println(x)
}