package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	money   int
}

func (p Person) Talk(msg string) string {
	return msg
}

// reflect.StructField结构的结构如下
/*
type StructField struct {
  // 字段名称
  Name string
  // 包名
  PkgPath string
  // 类型名
  Type      Type
  // Tag
  Tag       StructTag
  // 字段的字节偏移
  Offset    uintptr
  // 索引
  Index     []int
  // 是否为嵌套字段
  Anonymous bool
}
*/

func main() {
	// 访问结构体字段的方法有两种，一种是通过索引来进行访问，另一种是通过名称
	// 索引访问
	rType := reflect.TypeOf(new(Person)).Elem()
	// 输出结构体字段的数量
	fmt.Println(rType.NumField())
	for i := 0; i < rType.NumField(); i++ {
		structField := rType.Field(i)
		fmt.Println(structField.Index, structField.Name, structField.Type, structField.Offset, structField.IsExported())
		/*	输出：
			[0] Name string 0 true
			[1] Age int 16 true
			[2] Address string 24 true
			[3] money int 40 false (money是小写开头，所以是未导出的)
		*/
	}
	// 名称访问
	// 输出结构体字段的数量
	fmt.Println(rType.NumField())
	if field, ok := rType.FieldByName("money"); ok {
		fmt.Println(field.Name, field.Type, field.IsExported()) // money int false
	}

	// 倘若要修改结构体字段值，则必须传入一个结构体指针，下面是一个修改字段的例子
	// 传入指针
	rValue := reflect.ValueOf(&Person{
		Name:    "Dick",
		Age:     0,
		Address: "",
		money:   0,
	}).Elem()
	fmt.Println("原结构体：", rValue.Interface())

	// 获取字段
	name := rValue.FieldByName("Name")
	// 修改字段值
	if (name != reflect.Value{}) { // 如果返回reflect.Value{}，则说明该字段不存在
		name.SetString("Jack")
	}
	// 输出结构体
	fmt.Println("新结构体：", rValue.Interface())

	// 对于修改结构体私有字段而言，需要进行一些额外的操作
	// 传入指针
	rValue2 := reflect.ValueOf(&Person{
		Name:    "Fuck",
		Age:     0,
		Address: "",
		money:   0,
	}).Elem()
	fmt.Println("原结构体：", rValue2.Interface())

	// 获取一个私有字段
	money := rValue2.FieldByName("money")
	// 修改字段值
	if (money != reflect.Value{}) { // 有这个字段
		// 构造指向该结构体未导出字段的指针反射值
		p := reflect.NewAt(money.Type(), money.Addr().UnsafePointer()) // func reflect.NewAt(typ reflect.Type, p unsafe.Pointer) reflect.Value
		// 获取该指针所指向的元素，也就是要修改的字段
		field := p.Elem()
		// 修改值
		field.SetInt(164)
	}
	// 输出结构体
	fmt.Printf("新结构体：%+v\n", rValue2.Interface())

	/*
		访问方法与访问字段的过程很相似，只是函数签名略有区别。reflect.Method结构体如下
		type Method struct {
		  // 方法名
		  Name string
		  // 包名
		  PkgPath string
		  // 方法类型
		  Type  Type
		  // 方法对应的函数，第一个参数是接收者
		  Func  Value
		  // 索引
		  Index int
		}
	*/
	// 访问方法信息示例如下
	// 输出方法个数
	fmt.Println(rType.NumMethod())
	// 遍历输出方法信息
	for i := 0; i < rType.NumMethod(); i++ {
		method := rType.Method(i)
		fmt.Println(method.Index, method.Name, method.Type, method.IsExported())
		fmt.Println("方法参数")
		for i := 0; i < method.Func.Type().NumIn(); i++ {
			fmt.Println(method.Func.Type().In(i).String())
		}
		fmt.Println("方法返回值")
		for i := 0; i < method.Func.Type().NumOut(); i++ {
			fmt.Println(method.Func.Type().Out(i).String())
		}
	}

	/*
		调用方法：
		调用方法与调用函数的过程相似，而且并不需要手动传入接收者，例子如下
	*/
	// 遍历输出方法信息
	talk := rValue.MethodByName("Talk")
	if (talk != reflect.Value{}) {
		// 调用方法，并获取返回值
		res := talk.Call([]reflect.Value{reflect.ValueOf("hello,reflect!")})
		// 遍历输出返回值
		for _, re := range res {
			fmt.Println(re.Interface())
		}
	}

	/*
		创建：
		通过反射可以构造新的值，reflect包同时根据一些特殊的类型提供了不同的更为方便的函数
	*/
	/*
		基本类型：
		// 返回指向反射值的指针反射值
		func New(typ Type) Value
	*/
	// 以string为例
	rValue3 := reflect.New(reflect.TypeOf(*new(string)))
	rValue3.Elem().SetString("hello world!")
	fmt.Println(rValue3.Elem().Interface())
}
