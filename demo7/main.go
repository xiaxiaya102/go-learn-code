package main

import (
	"fmt"
	"strings"
)

//func main() {
//	//scoreMap := make(map[string]int, 8)
//	//scoreMap["张三"] = 90
//	//scoreMap["小明"] = 100
//	//fmt.Println(scoreMap)
//	//fmt.Println(scoreMap["小明"])
//	//fmt.Printf("type of a:%T\n", scoreMap)
//
//	userInfo := map[string]string{
//		"username": "zhangsan",
//		"password": "123456",
//	}
//
//	userInfo["age"] = "178"
//
//	for k, v := range userInfo {
//		fmt.Println(k, v)
//	}
//
//}

//func main() {
//	scoreMap := make(map[string]int)
//	scoreMap["张三"] = 90
//	scoreMap["小明"] = 100
//	// 如果 key 存在 ok 为 true,v 为对应的值；不存在 ok 为 false,v 为值类型的零值
//	v, ok := scoreMap["张三1"]
//	if ok {
//		fmt.Println(v)
//	} else {
//		fmt.Println(v)
//		fmt.Println("查无此人")
//	}
//}

//func main() {
//	//map 的遍历
//	scoreMap := make(map[string]int)
//	scoreMap["张三"] = 90
//	scoreMap["小明"] = 100
//	scoreMap["娜扎"] = 60
//	for _, v := range scoreMap {
//		fmt.Println(v)
//	}
//}

//func main() {
//	scoreMap := make(map[string]int)
//	scoreMap["张三"] = 90
//	scoreMap["小明"] = 100
//	scoreMap["娜扎"] = 60
//	delete(scoreMap, "小明") //将小明:100 从 map 中删除
//	for k, v := range scoreMap {
//		fmt.Println(k, v)
//	}
//}
//
//func main() {
//	rand.NewSource(time.Now().UnixNano()) //初始化随机数种子
//	var scoreMap = make(map[string]int, 200)
//	for i := 0; i < 100; i++ {
//		key := fmt.Sprintf("stu%02d", i) //生成 stu 开头的字符串
//		value := rand.Intn(100)          //生成 0~99 的随机整数
//		scoreMap[key] = value
//	}
//	//取出 map 中的所有 key 存入切片 keys
//	var keys = make([]string, 0, 200)
//	for key := range scoreMap {
//		keys = append(keys, key)
//	}
//	//对切片进行排序
//	sort.Strings(keys)
//	//按照排序后的 key 遍历 map
//	for _, key := range keys {
//		fmt.Println(key, scoreMap[key])
//	}
//}

//func main() {
//	var mapSlice = make([]map[string]string, 3)
//	for index, value := range mapSlice {
//		fmt.Printf("index:%d value:%v\n", index, value)
//	}
//	fmt.Println("after init")
//	// 对切片中的 map 元素进行初始化
//	mapSlice[0] = make(map[string]string, 10)
//	mapSlice[0]["name"] = "小王子"
//	mapSlice[0]["password"] = "123456"
//	mapSlice[0]["address"] = "海淀区"
//	for index, value := range mapSlice {
//		fmt.Printf("index:%d value:%v\n", index, value)
//	}
//}

//func main() {
//	var sliceMap = make(map[string][]string, 3)
//	fmt.Println(sliceMap)
//	fmt.Println("after init")
//	key := "中国"
//	value, ok := sliceMap[key]
//	fmt.Println(value, ok)
//	if !ok {
//		value = make([]string, 0, 2)
//		fmt.Println(value)
//	}
//	value = append(value, "北京", "上海")
//	sliceMap[key] = value
//	fmt.Println(sliceMap)
//}

func main() {
	/**
	practice:写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中 how=1 do=2
	you=1。
	*/
	s := "how do you do are you ok"
	strSlice := strings.Split(s, " ")
	resultMap := make(map[string]int, 4)
	fmt.Println(resultMap)
	for _, word := range strSlice {
		resultMap[word]++
	}
	fmt.Println(resultMap)
}
