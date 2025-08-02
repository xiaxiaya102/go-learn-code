package main

import "fmt"

var m = map[string]int{"one": 1, "two": 2, "three": 3}

func main() {

	//常用变量
	// 布尔型
	//var isGoFun bool = true
	//fmt.Println("bool")
	//fmt.Println(isGoFun)
	//
	//// 整型
	//var a int = 10
	//var b uint = 20
	//fmt.Println("int")
	//fmt.Println(a + int(b))
	//
	//// 浮点型
	//var pi float64 = 3.14159
	//fmt.Println("float")
	//fmt.Println(pi)
	//
	//// 字符串
	//var hello string = "Hello, Go!"
	//fmt.Println("string")
	//fmt.Println(hello)
	//
	//// 数组
	//var arr [3]int = [3]int{1, 2, 3}
	//fmt.Println("array")
	//fmt.Println(arr)
	//
	//// 切片
	//var slice []int = []int{1, 2, 3, 4}
	//fmt.Println("slice")
	//fmt.Println(slice)
	//
	//// 映射
	//var m map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}
	//fmt.Println("map")
	//fmt.Println(m)
	//
	//// 结构体
	//type Point struct {
	//	x int
	//	y int
	//}
	//fmt.Println("struct")
	//p := Point{x: 3, y: 4}
	//fmt.Println(p)
	//
	//// 指针
	//var ptr *int = &a
	//fmt.Println("pointer")
	//fmt.Println(*ptr)
	//
	//// 接口
	//var i interface{} = "Hello"
	//fmt.Println("interface")
	//fmt.Println(i)
	//
	//// 函数
	//add := func(x, y int) int {
	//	return x + y
	//}
	//fmt.Println("function")
	//fmt.Println(add(5, 3))

	//批量声明变量并赋值
	//var (
	//	username  string  = "gopher"
	//	age       int     = 18
	//	isStudent bool    = true
	//	height    float64 = 1.8
	//)
	//
	//fmt.Println("username:", username, "age:", age, "isStudent:", isStudent, "height:", height)
	//
	//var (
	//	a string
	//	b int
	//	c bool
	//)
	//a = "张三"
	//b = 10
	//c = true
	//fmt.Println(a, b, c)
	//
	//n := 1
	//for i := 1; i <= 100; i++ {
	//	n = n * i
	//}

	//	_, username := getInfo()
	//	fmt.Println(username)
	//
	//	fmt.Println(pi)
	//
	//	const pi = math.Pi
	//	fmt.Println(pi)
	//
	//	num8 := 5.1234e2   // ? 5.1234 * 10 的 2 次方
	//	num9 := 5.1234e2   // ? 5.1234 * 10 的 2 次方 shift+alt+向下的箭头
	//	num10 := 5.1234e-2 // ? 5.1234 / 10 的 2 次方 0.051234
	//	fmt.Println("num8=", num8, "num9=", num9, "num10=", num10)
	//
	//	fmt.Println("str := \"c:\\Code\\demo\\go.exe\"")
	//
	//	s1 := `第一行
	//第二行
	//第三行`
	//	fmt.Println(strings.Split(s1, "\n"))

	//var str = "this is golang"
	//var flag = strings.Contains(str, "golang")
	//fmt.Println(flag)

	//var str = "this is golang"
	//var flag = strings.HasSuffix(str, "golang")
	//fmt.Println(flag)

	//var str = "123-456-789"
	//var arr = strings.Split(str, "-")
	//fmt.Println(arr)
	//var str2 = strings.Join(arr, "*")
	//fmt.Println(str2)
	//

	changeString()
}

func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))
	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

const (
	pi = 3.1415
	e  = 2.7182
)

const (
	n1 = 100
	n2
	n3
)

func getInfo() (int, string) {
	return 10, "张三"
}
