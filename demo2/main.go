package main

func main() {
	//fmt.Println("类型转换")
	//var a int = 10
	//var b float32 = 20.9
	//var c = a + int(b)
	//fmt.Printf("c的类型是：%T,c的值为：%d\n", c, c)
	//var d = float32(a) + b
	//fmt.Printf("d的类型是：%T,d的值为：%f\n", d, d)
	//var e float64 = 30.9
	//fmt.Printf("e的类型是：%T,e的值为：%f\n\n", e, e)
	//var f = float32(e) + b
	//fmt.Printf("f的类型是：%T,f的值为：%f\n", f, f)

	//var a int16 = 129
	//var b = int8(a)  // 范围 -128 到 127
	//println("b=", b) //b= -127 //错误

	//var a, b = 3, 4
	//var c int
	//// math.Sqrt()接收的参数是 float64 类型，需要强制转换
	//c = int(math.Sqrt(float64(a*a + b*b)))
	//fmt.Println(c)

	//其他类型转换成 String 类型
	//var i int = 20
	//var f float64 = 12.456
	//var t bool = true
	//var b byte = 'a'
	//var strs string
	//strs = fmt.Sprintf("%d", i)
	//fmt.Printf("str type %T ,strs=%v \n", strs, strs)
	//strs = fmt.Sprintf("%f", f)
	//fmt.Printf("str type %T ,strs=%v \n", strs, strs)
	//strs = fmt.Sprintf("%t", t)
	//fmt.Printf("str type %T ,strs=%v \n", strs, strs)
	//strs = fmt.Sprintf("%c", b)
	//fmt.Printf("str type %T ,strs=%v \n", strs, strs)

	//使用 strconv 包里面的几种转换方法进行转换
	//1、int 转换成 string
	//var num1 int = 20
	//s1 := strconv.Itoa(num1)
	//fmt.Printf("str type %T ,strs=%v \n", s1, s1)
	//// 2、float 转 string
	//var num2 float64 = 20.113123
	///* 参数 1：要转换的值
	//   参数 2：格式化类型
	//   'f'（-ddd.dddd）、
	//   'b'（-ddddp±ddd，指数为二进制）、
	//   'e'（-d.dddde±dd，十进制指数）、
	//   'E'（-d.ddddE±dd，十进制指数）、
	//   'g'（指数很大时用'e'格式，否则'f'格式）、
	//   'G'（指数很大时用'E'格式，否则'f'格式）。
	//   参数 3: 保留的小数点 -1（不对小数点格式化）
	//   参数 4：格式化的类型
	//*/
	//s2 := strconv.FormatFloat(num2, 'f', 2, 64)
	//fmt.Printf("str type %T ,strs=%v \n", s2, s2)
	//// 3、bool 转 string
	//s3 := strconv.FormatBool(true)
	//fmt.Printf("str type %T ,strs=%v \n", s3, s3)
	////4、int64 转 string
	//var num3 int64 = 20
	///*第二个参数为 进制
	// */
	//s4 := strconv.FormatInt(num3, 10)
	//fmt.Printf("类型 %T ,strs=%v \n", s4, s4)

	//var s = "123456789"
	//i64, _ := strconv.ParseInt(s, 10, 64)
	//fmt.Printf("值：%v 类型：%T", i64, i64)

	// string 转字符
	//s := "hello 张三"
	//for _, r := range s { //rune
	//	fmt.Printf("%v(%c) ", r, r)
	//}
	//fmt.Println()
	//
	//for i := 0; i < len(s); i++ {
	//	fmt.Printf("%v(%c) ", s[i], s[i])
	//}

}
