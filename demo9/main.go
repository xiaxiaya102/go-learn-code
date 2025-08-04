package main

import (
	"fmt"
	"time"
)

func unixToTime(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	year := timeObj.Year()             //年
	month := timeObj.Month()           //月
	day := timeObj.Day()               //日
	hour := timeObj.Hour()             //小时
	minute := timeObj.Minute()         //分钟
	second := timeObj.Second()         //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

//func main() {
//	unixToTime(1587880013)
//	t := time.Now()
//
//	later := t.Add(time.Hour)
//	laterStr := later.Format("2006-01-02 15:04:05")
//	fmt.Println(laterStr)
//
//	before := t.Sub(later)
//	fmt.Println(before)
//
//}

func main() {
	// 1、获取当前时间，格式化输出为 2020/06/19 20:30:05 格式。
	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05"))

	// 2、获取当前时间，格式化输出为时间戳
	timestrap := now.Unix()
	fmt.Printf("当前时间戳为：%d\n", timestrap)

	// 3、把时间戳 1587880013 转换成日期字符串
	nowTime := time.Unix(1587880013, 0) // 只传秒，纳秒传0
	fmt.Println(nowTime.Format("2006/01/02 15:04:05"))

	// 4、把日期字符串 2020/06/19 20:30:05 转换成时间戳
	const shortForm = "2006/01/02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation(shortForm, "2020/06/19 20:30:05", loc)
	fmt.Println(t)

	// 5、编写程序统计一段代码的执行耗时时间，单位精确到微秒。
	costTime(1000, 1000)
}

func costTime(x, y int) int {
	startTime := time.Now()
	sum := (x + y) * 2
	endTime := time.Now()
	fmt.Printf("costTime方法耗时：%d微秒\n", endTime.Sub(startTime)/time.Microsecond) // 转换成微秒
	return sum
}
