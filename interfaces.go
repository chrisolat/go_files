package main
import "fmt"

type Dog struct{
	age int
	speak string
}

type Cat struct{
	age int
	speak string
}

func (d* Dog) say(times int) {
	for range times {
		fmt.Println(d.speak)
	}
}

func (c* Cat) say(times int) {
	for range times {
		fmt.Println(c.speak)
	}
}

func (c* Cat) getage() int {
	return c.age
}

type animal interface {
	say(times int)
	getage() int
}

func say_smth(a animal, times int){
	a.say(times)
	fmt.Println(a.getage())
}

func main() {
	cat := &Cat{age:2, speak:"Meow"}
	//dog := &Dog{age:5, speak:"bark"}
	say_smth(cat, 2)
	//say_smth(dog, 3)
}
