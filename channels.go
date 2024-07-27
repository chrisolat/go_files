package main

import "fmt"

func main() {
	ch := make(chan int)
	go remove(ch)
	ch <- 1
}

func remove(ch chan int) {
	v := <- ch
	fmt.Println(v)
}