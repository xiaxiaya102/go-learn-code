package main

//type Person struct {
//	Name string
//	Age  int
//	Sex  string
//}
//
//func (p Person) getInfo() string {
//	return "姓名：" + p.Name + "，年龄：" + strconv.Itoa(p.Age) + ",性别：" + p.Sex
//}
//
//func main() {
//	//第一种实例化方式
//	var p1 Person
//	p1.Name = "xufulai"
//	p1.Age = 24
//	p1.Sex = "男"
//
//	//第二种 new 关键字对结构体进行实例化
//	var p2 = new(Person)
//	p2.Name = "xufulai"
//	p2.Age = 24
//	p2.Sex = "男"
//
//	//使用&对结构体进行取地址操作相当于对该结构体类型进行了一次 new 实例化操作
//	var p3 = &Person{}
//	p3.Name = "xufulai"
//	p3.Age = 24
//	p3.Sex = "男"
//
//	//结构体实例化（第四种方法） 键值对初始化
//	p4 := Person{
//		Name: "xufulai",
//		Age:  24,
//		Sex:  "男",
//	}
//	fmt.Println(p4.getInfo())
//
//	//结构体实例化（第四种方法） 键值对初始化
//	p5 := &Person{
//		Name: "xufulai",
//		Age:  24,
//		Sex:  "男",
//	}
//
//	fmt.Println(p5.getInfo())
//}

type Animal struct {
	Name string
	Age  int
}

type Dog struct {
}

func main() {

}
