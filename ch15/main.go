package main

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
)

/*
*
运行时反射：字符串和结构体之间如何转换？
*/
func main() {
	//reflectTest1()
	//reflectTest2()
	//loopPersonAttribute()
	//isImpMethod()
	//jsonToStruct()
	//jsonToStructWithTag()
	reflectCallMethod()
}

// reflect.Value 和 reflect.Type 来分别表示变量的值和类型
func reflectTest1() {
	i := 3
	iv := reflect.ValueOf(i)
	it := reflect.TypeOf(i)
	fmt.Println(iv, it)

	// reflect.Value to int
	i2 := iv.Interface().(int)
	fmt.Println(i2)

	// 反射修改值, 需要使用指针
	ipv := reflect.ValueOf(&i)
	ipv.Elem().SetInt(4)
	fmt.Println(i)

}

// 结构体属性修改
func reflectTest2() {
	p := person{Name: "飞雪无情", Age: 20}
	pv := reflect.ValueOf(p)
	ppv := reflect.ValueOf(&p)
	ppv.Elem().FieldByName("Name").SetString("张三")
	fmt.Println(p)
	fmt.Println("ppv type:", ppv.Kind())
	fmt.Println("pv type:", pv.Kind())
}

type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("Name is %s,Age is %d", p.Name, p.Age)
}

// 反射遍历person属性和方法
func loopPersonAttribute() {
	p := person{Name: "飞雪无情", Age: 20}
	pt := reflect.TypeOf(p)
	//遍历person的字段
	for i := 0; i < pt.NumField(); i++ {
		fmt.Println("字段：", pt.Field(i).Name)
	}
	//遍历person的方法
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Println("方法：", pt.Method(i).Name)
	}
}

// 是否实现指定方法
func isImpMethod() {
	p := person{Name: "飞雪无情", Age: 20}
	pt := reflect.TypeOf(p)
	stringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	fmt.Println("是否实现了fmt.Stringer：", pt.Implements(stringerType))
	fmt.Println("是否实现了io.Writer：", pt.Implements(writerType))
}

// json to struct
func jsonToStruct() {
	p := person{Name: "飞雪无情", Age: 20}
	// struct to json
	jsonB, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(jsonB))
	}
	// json to struct
	respJSON := "{\"Name\":\"李四\",\"Age\":40}"

	p2 := person{}
	err = json.Unmarshal([]byte(respJSON), &p2)
	if err != nil {
		return
	}
	fmt.Println(p2)
}

type personWithTag struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 使用tag指定json的key
func jsonToStructWithTag() {
	p := personWithTag{Name: "飞雪无情", Age: 20}
	// struct to json
	jsonB, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(jsonB))
	}

	fmt.Println()

	pt := reflect.TypeOf(p)
	//遍历person字段中key为json的tag
	for i := 0; i < pt.NumField(); i++ {
		sf := pt.Field(i)
		fmt.Printf("字段%s上,json tag为%s\n", sf.Name, sf.Tag.Get("json"))
	}
}

func (p person) Print(prefix string) {
	fmt.Printf("%s:Name is %s,Age is %d\n", prefix, p.Name, p.Age)
}

// 反射调用方法
func reflectCallMethod() {
	p := person{Name: "飞雪无情", Age: 20}
	pv := reflect.ValueOf(p)
	//反射调用person的Print方法
	mPrint := pv.MethodByName("Print")
	args := []reflect.Value{reflect.ValueOf("登录")}
	mPrint.Call(args)
}
