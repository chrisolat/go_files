package main

import "fmt"

type test struct{
	val int
	ch rune
}

func get_test() *test {
	return &test{
		val: 2,
		ch: 'b',
	}
}

func (t test) addval(i int){
	t.val += i
}

func main(){
	fmt.Println("Hello")
	tst := test{val:1,ch:'a'}
	var tst2 *test = get_test()
	(*tst2).val = 3
	fmt.Printf("%+v\n",tst)
	fmt.Printf("%d\n",(*tst2).val)
	
	var tst3* test
	tst3 = &test{1,'b'}
	fmt.Printf("%+v\n",tst3)
	tst3.addval(2)
	fmt.Printf("%+v\n",tst3)
}