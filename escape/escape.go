package main

// 1. 返回了变量的指针造成外部引用
func quote() (*int, []int, map[int]int) {
	i := 0                        // moved to heap: i
	list := []int{1, 2, 3, 4}     // 底层数组 []int{...} escapes to heap
	mp := map[int]int{1: 1, 2: 2} // 底层hmap结构体 map[int]int{...} escapes to heap

	return &i, list, mp
}

func main() {
	quote()
}
