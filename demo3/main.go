package main

import "fmt"

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

	for i := 0; i < 10; i++ {
		fmt.Println(i)
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
