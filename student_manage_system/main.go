package main

import (
	"fmt"
	"sort"
	"strings"
)

// 定义一个切片,模拟数据库
var database = make([]Student, 0)

// Student 学生信息结构体
type Student struct {
	Name  string
	Age   int
	Score float64
}

// GetInfo 实现 StudentInfo 接口的 GetInfo 方法
func (s Student) GetInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d, Score: %.2f", s.Name, s.Age, s.Score)
}

// AddStudent 增加学生信息
func AddStudent(student Student) {
	database = append(database, student)
}

// DeleteStudentByName 根据名称删除学生信息
func DeleteStudentByName(name string) {
	for i, student := range database {
		if student.Name == name {
			database = append(database[0:i], database[i+1:]...)
		}
	}
}

// UpdateStudent 根据学生名称更新学生信息
func UpdateStudent(name string, age int, score float64) {
	for i, oldStudentInfo := range database {
		if oldStudentInfo.Name == name {
			var newStudent = Student{
				name,
				age,
				score,
			}
			database[i] = newStudent
		}
	}
}

// FindStudent 根据学生姓名查找学生信息
func FindStudent(name string) Student {
	for i, studentInfo := range database {
		if studentInfo.Name == name {
			return database[i]
		}
	}
	return Student{}
}

// 实现 sort.Interface 接口
type ByScore []Student

func (s ByScore) Len() int           { return len(s) }
func (s ByScore) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByScore) Less(i, j int) bool { return s[i].Score > s[j].Score } // 从高到低排序

// ListStudentByScore 根据学生分数排序，从高到低
func ListStudentByScore(studentList []Student) []Student {
	sort.Sort(ByScore(studentList))
	return studentList
}

// ListStudentByAge 根据学生年龄排序，从低到高
func ListStudentByAge(studentList []Student) []Student {
	sort.Slice(studentList, func(i, j int) bool {
		return studentList[i].Age < studentList[j].Age
	})

	return studentList
}

// CalculateAverageScoreAndPassNum 计算所有学生平均成绩和通过考试的人数
func CalculateAverageScoreAndPassNum(studentList []Student) (float64, int) {
	scoreSum := 0.0
	passSum := 0
	for _, student := range studentList {
		scoreSum = +student.Score
		if student.Score >= 60 {
			passSum++
		}
	}

	return scoreSum / float64(len(studentList)), passSum
}

// checkName 检查学生姓名是否包含指定的子字符串
func checkName(studentList []Student, tag string) []Student {
	newStudent := make([]Student, 0)
	for i, student := range studentList {
		if strings.Contains(student.Name, tag) {
			newStudent = append(newStudent, studentList[i])
		}
	}
	return newStudent
}

// upperName 将所有学生的姓名转换为大写
func upperName(studentList *[]Student) *[]Student {
	for _, student := range *studentList {
		student.Name = strings.ToUpper(student.Name)
	}
	return studentList
}

// StudentInfo 定义接口
type StudentInfo interface {
	GetInfo() string
}

func main() {
	database := []Student{
		{"张三", 20, 88.5},
		{"李四", 22, 92.0},
		{"王五", 21, 75.0},
		{"赵六", 23, 85.5},
		{"孙七", 24, 59.5},
	}

	var student = Student{
		"chensimei",
		22,
		99.0,
	}

	// 添加学生
	AddStudent(student)

	// 删除学生
	DeleteStudentByName("xufulai")

	fmt.Println("删除学生后：", database)

	// 修改成绩和年龄
	UpdateStudent("chensimei", 23, 100.0)

	fmt.Println("更新后学生后：", database)

	// 根据名称查询学生信息
	studentInfo := FindStudent("chensimei")
	fmt.Println("陈思梅的学生信息是：", studentInfo)

	//成绩高低排序
	studentListByScoreDesc := ListStudentByScore(database)
	fmt.Println("成绩由高到低是：", studentListByScoreDesc)

	//年龄低到高排序
	studentListByAge := ListStudentByAge(database)
	fmt.Println("年龄由低到高是：", studentListByAge)

	//计算平均成绩 统计通过人数
	averageScore, passNum := CalculateAverageScoreAndPassNum(database)
	fmt.Println("平均成绩：", averageScore, "考试通过人数：", passNum)

	//检查学生姓名是否包含指定的子字符串
	findResult := checkName(database, "x")
	fmt.Println("名字中包含x的学生信息：", findResult)

	//将所有学生的姓名转换为大写
	upperResult := upperName(&database)
	fmt.Println("姓名转化为大写后的学生信息：", upperResult)

	//循环控制 continue
	for _, student := range database {
		if student.Score < 60 {
			continue
		}
		fmt.Println("成绩合格：", student)
	}

	//循环控制 break
	for _, student := range database {
		if student.Score < 60 {
			break
		}
		fmt.Println("成绩合格：", student)
	}

	//空接口

	// 创建一些不同类型的变量
	num := 100
	str := "Hello, Go!"
	userInfo := map[string]interface{}{
		"name": "xfl",
		"age":  24,
		"sex":  "男",
	}

	// 定义一个空接口变量，可以存储任何类型
	var emptyInterface interface{}

	// 将不同类型的数据存储到空接口中
	emptyInterface = student
	fmt.Println("存储 Student 类型：", emptyInterface)
	if v, ok := emptyInterface.(Student); ok {
		fmt.Println(v)
	}

	emptyInterface = num
	fmt.Println("存储 int 类型：", emptyInterface)

	emptyInterface = str
	fmt.Println("存储 string 类型：", emptyInterface)

	// 使用类型断言从空接口中提取具体类型
	if s, ok := emptyInterface.(Student); ok {
		fmt.Println("类型断言成功，Student 类型：", s)
	}

	if n, ok := emptyInterface.(int); ok {
		fmt.Println("类型断言成功，int 类型：", n)
	}

	t, ok := emptyInterface.(string)
	if ok {
		fmt.Println("类型断言成功，string 类型：", t)
	}

	emptyInterface = userInfo
	if u, ok := emptyInterface.(map[string]string); ok {
		fmt.Println(u)
	}

	// 使用接口调用 GetInfo 方法
	var info StudentInfo = student
	fmt.Println(info.GetInfo()) // 输出：Name: 张三, Age: 20, Score: 88.50

	//统计各分数段人数
	sourceMap := make(map[string]int)
	for _, s := range database {
		if s.Score > 90 {
			sourceMap["90+"]++
		} else if s.Score > 80 && s.Score < 90 {
			sourceMap["80+"]++
		} else if s.Score > 70 && s.Score < 80 {
			sourceMap["70+"]++
		} else if s.Score > 60 && s.Score < 70 {
			sourceMap["60+"]++
		}
	}

	fmt.Println("各分段成绩为:", sourceMap)
	for score, num := range sourceMap {
		fmt.Printf("分数段: %s 人数为: %d", score, num)
	}

	if count, ok := sourceMap["80+"]; ok {
		fmt.Printf("\n80+分数段的学生人数: %d\n", count)
	}

}
