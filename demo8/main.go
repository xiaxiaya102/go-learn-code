package main

import (
	"errors"
	"fmt"
)

//func main() {
//var result int
//result = intSum(1, 3)
//fmt.Println(result)
//sayHello()
//ret := intSum3(1, 2, 3, 4)
//fmt.Println(ret)
//intSum4(1, 2, 3, 4, 5, 6)
//sub, sum := calc(1, 2)
//fmt.Println(sub, sum)
//
//sum, sub := calc(1, 2)
//fmt.Println(sum, sub)
//}

func intSum(x int, y int) int {
	return x + y
}

func intSum2(x, y int) int {
	return x + y
}

func intSum3(x ...int) int {
	fmt.Printf("x的类型：%T", x)
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

// 固定参数搭配可变参数使用时，可变参数要放在固定参数的后面
func intSum4(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

func sayHello() {
	fmt.Println("Hello world")
}

// 返回值命名
func calc1(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return sub, sum
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func do(s string) func(int, int) int {
	switch s {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

type calculation func(int, int) int

//func main() {
//	var c calculation               // 声明一个 calculation 类型的变量 c
//	c = add                         // 把 add 赋值给 c
//	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
//	fmt.Println(c(1, 2))            // 像调用 add 一样调用 c
//	f := add                        // 将函数 add 赋值给变量 f1
//	fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
//	fmt.Println(f(10, 20))          // 像调用 add 一样调用 f
//	d := sayHello
//	fmt.Printf("type of d:%T\n", d)
//}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

//func main() {
//	// 函数作为参数传入
//	ret2 := calc(10, 20, add)
//	fmt.Println(ret2) //30
//}
//
//func main() {
//	var a = do("+")
//	fmt.Println(a(10, 20))
//}

//func main() {
//	// 将匿名函数保存到变量
//	add := func(x, y int) {
//		fmt.Println(x + y)
//	}
//	add(10, 20) // 通过变量调用匿名函数
//
//	//自执行函数：匿名函数定义完加()直接执行
//	func(x, y int) {
//		fmt.Println(x + y)
//	}(10, 20)
//}

//func adder() func(int) int {
//	var x int
//	return func(y int) int {
//		x += y
//		return x
//	}
//}
//
//func main() {
//	var f = adder()
//	fmt.Println(f(10)) //10
//	fmt.Println(f(20)) //30
//	fmt.Println(f(30)) //60
//	f1 := adder()
//	fmt.Println(f1(40)) //40
//	fmt.Println(f1(50)) //90
//}

//func makeSuffixFunc(suffix string) func(string) string {
//	return func(name string) string {
//		if !strings.HasSuffix(name, suffix) {
//			return name + suffix
//		}
//		return name
//	}
//}
//
//func main() {
//	jpgFunc := makeSuffixFunc(".jpg")
//	txtFunc := makeSuffixFunc(".txt")
//	fmt.Println(jpgFunc("test")) //test.jpg
//	fmt.Println(txtFunc("test")) //test.txt
//}

//func f1() int {
//	x := 5
//	defer func() {
//		// 已经确定了x的值了 5
//		x++
//	}()
//	return x
//}
//func f2() (x int) {
//	//具名返回值,指定了返回x 6
//	defer func() {
//		x++
//	}()
//	return 5
//}
//func f3() (y int) {
//	//具名返回值,指定了返回y
//	x := 5
//	defer func() {
//		//返回前修改x值不会影响y值  return 5
//		x++
//	}()
//	// y 值 5
//	return x
//}
//func f4() (x int) {
//	//具名返回值x
//	defer func(x int) {
//		//（x） 值引用,相当于复制了一个变量，不影响原来x的值
//		x++
//	}(x)
//	// x-> 5
//	return 5
//}
//
//func main() {
//	fmt.Println(f1()) //5
//	fmt.Println(f2()) //6
//	fmt.Println(f3()) //5
//	fmt.Println(f4()) //5
//}

// panic/recover 的基本使用
//func funcA() {
//	fmt.Println("func A")
//}
//func funcB() {
//	panic("panic in B")
//}
//func funcC() {
//	fmt.Println("func C")
//}
//func main() {
//	funcA()
//	funcB() //panic 导致程序奔溃退出，无法执行funcC
//	funcC()
//}

//func funcA() {
//	fmt.Println("func A")
//}
//func funcB() {
//	//recover()必须搭配 defer 使用。
//	defer func() {
//		err := recover()
//		fmt.Println(err)
//		//如果程序出出现了 panic 错误,可以通过 recover 恢复过来
//		if err != nil {
//			fmt.Println("recover in B")
//		}
//	}()
//	panic("panic in B")
//}
//func funcC() {
//	fmt.Println("func C")
//}
//func main() {
//	funcA()
//	funcB()
//	funcC()
//}

//func fn2() {
//	defer func() {
//		err := recover()
//		if err != nil {
//			fmt.Println("抛出异常给管理员发送邮件")
//			fmt.Println(err)
//		}
//	}()
//	num1 := 10
//	num2 := 0
//	res := num1 / num2
//	fmt.Println("res=", res)
//}
//
//func main() {
//	fn2()
//}

func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	}
	return errors.New("读取文件错误")
}
func fn3() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("抛出异常给管理员发送邮件")
		}
	}()
	var err = readFile("main.go")
	if err != nil {
		panic(err)
	}
	fmt.Println("继续执行")
}
func main() {
	fn3()
}
