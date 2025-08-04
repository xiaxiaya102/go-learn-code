package main

import "fmt"

//func main() {
//	var a = 10
//	var b = &a //取地址操作
//	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc0000100a8
//	fmt.Printf("b:%v type:%T\n", b, b) // b:0xc0000100a8 type:*int
//	fmt.Println("取 b 的地址：", &b)        // 0xc000006028
//}

//func main() {
//	a := 10
//	b := &a // 取变量 a 的地址，将地址保存到指针 b 中
//	fmt.Printf("type of b:%T\n", b)
//	c := *b // 指针取值（根据指针的值去内存取值）
//	fmt.Printf("type of c:%T\n", c)
//	fmt.Printf("value of c:%v\n", c)
//}

func main() {
	userinfo := map[string]string{}

	userinfo["username"] = "张三"
	fmt.Println(userinfo)
}

func main1() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false
}

func main2() {
	a := 10
	b := &a // b 是 a 的指针

	fmt.Println("a 的值:", a)    // 输出 a 的值，10
	fmt.Println("b 的地址:", b)   // 输出 b 的地址，a 的地址
	fmt.Println("b 指向的值:", *b) // 输出 b 指向的值，10

	*b = 20                  // 通过指针修改 a 的值
	fmt.Println("a 的新值:", a) // 输出 a 的新值，20
}
