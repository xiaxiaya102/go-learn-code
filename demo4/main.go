package main

import (
	"fmt"
	"strings"
)

func main() {
	//var name string
	//fmt.Println("请输入姓名：")
	//fmt.Scanln(&name)
	//
	//fmt.Println("请输入年龄：")
	//var age int
	//fmt.Scanln(&age)
	//
	//fmt.Println("请输入成绩：")
	//var score float64
	//fmt.Scanln(&score)
	//
	//// 调用 analyzeScore 函数并将 score 作为参数传入
	//fmt.Printf("您好，姓名：%s\n您本次的成绩：%.2f\n", name, score)
	//// 调用 analyzeScore 函数来分析成绩
	//result := analyzeScore(score)
	//fmt.Println("成绩评定：", result)

	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}

	//str := "abc 上海"
	//for index, val := range str {
	//	fmt.Printf("index=%d, val=%c \n", index, val)
	//}
	////
	//for _, val := range str {
	//	fmt.Printf("val=%c \n", val)
	//}
	//fileName := "hello.txt"
	//fmt.Printf("文件：%s 是：%s 类型文件", fileName, judgeFireType(fileName))

	//switchDemo5()

	//k := 1
	//for { // 这里也等价 for ; ; {
	//	if k <= 10 {
	//		fmt.Println("ok~~", k)
	//	} else {
	//		break //break 就是跳出这个 for 循环
	//	}
	//	k++
	//}

label2:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break label2
			}
			fmt.Println("i j 的值", i, "-", j)
		}
	}
}

// analyzeScore 函数，用于根据成绩返回评定结果
func analyzeScore(score float64) string {
	if score >= 60 {
		return "恭喜通过"
	} else {
		return "请重新考试"
	}
}

func judgeFireType(fileName string) string {
	fileExt := strings.Split(fileName, ".")
	switch fileExt[len(fileExt)-1] {
	case "txt":
		return "文本文件"
	case "html":
		return "网页文件"
	case "javascript":
		return "js脚本文件"
	case "css":
		return "样式文件"

	default:
		return ""
	}

}

func switchDemo5() {
	s := "c"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "c":
		fmt.Println("2")
	case s == "b":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
