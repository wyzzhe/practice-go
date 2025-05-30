package main

import "fmt"

type animal interface {
	run()
}

type dog struct{}

func (d dog) run() {}

// 1. 返回了变量的指针造成外部引用
func quote() (*int, []int, map[int]int) {
	i := 0                        // moved to heap: i
	list := []int{1, 2, 3, 4}     // 底层数组 []int{...} escapes to heap
	mp := map[int]int{1: 1, 2: 2} // 底层hmap结构体 map[int]int{...} escapes to heap

	return &i, list, mp
}

// 2. 把变量的指针传入通道可能会造成外部引用 （保守）
func channel() {
	i := 2 // i 逃逸，通道里指向i的指针可能会被外部引用
	ch := make(chan *int, 2)
	ch <- &i
	<-ch
}

// 3. 变量被闭包捕获并返回造成外部引用
func cluster() func() {
	i := 1
	return func() {
		fmt.Println(i)
	}
}

// 4. 切片存储变量地址，变量被切片间接引用导致逃逸 （保守）
func sliceMap() {
	i := 1
	list := make([]*int, 10)
	list[0] = &i
}

// 5. 接口变量可能被外部引用导致逃逸 （保守）
func interfaceEscape() {
	a := dog{}
	a.run()

	var aE animal
	aE = dog{}
	aE.run()
}

func main() {
	quote()
	channel()
	cluster()
	sliceMap()
	interfaceEscape()
}
