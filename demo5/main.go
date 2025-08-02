package main

import (
	"fmt"
	"sort"
)

func main() {
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	sort.Ints(intList)
	fmt.Println(intList)
	//int数组转int切片
	intSlice := sort.IntSlice(intList)
	fmt.Println(intSlice)
	sort.Sort(sort.Reverse(intSlice))
	fmt.Println(intList)
}
