package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

/*
	type Type interface{
		Kind() Kind
		Elem() Type
		Size() uintptr
		Elem() Type
		Implements(u Type) bool
		ConvertibleTo(u Type) bool
		......
	}
*/

func main() {
	// reflect.Type代表着 Go 中的类型，使用reflect.TypeOf()函数可以将变量转换成reflect.Type
	str := "Hello World!"
	reflectType := reflect.TypeOf(str)
	fmt.Println(reflectType) // string

	// 通过Kind，可以知晓空接口存储的值究竟是什么基础类型
	var eface any
	eface = 114514
	fmt.Println(reflect.TypeOf(eface).Kind()) // int

	// 使用Type.Elem()方法，可以判断类型为any的数据结构所存储的元素类型，可接收的底层参数类型必须是指针，切片，数组，通道，映射表其中之一，否则会panic
	var eface2 any = map[string]int{}
	rType := reflect.TypeOf(eface2)
	fmt.Println(rType.Kind())        // map
	fmt.Println(rType.Key().Kind())  // string (key() 会返回map的键反射类型)
	fmt.Println(rType.Elem().Kind()) // int

	// 对于指针使用Elem()会获得其指向元素的反射类型
	var eface3 any
	eface3 = new(strings.Builder)
	rType2 := reflect.TypeOf(eface3)
	vType3 := rType2.Elem()
	// 输出包路径
	fmt.Println(vType3.PkgPath()) // strings
	// 输出类型名
	fmt.Println(vType3.Name()) // Builder

	// 通过Size方法可以获取对应类型所占的字节大小
	fmt.Println(reflect.TypeOf(123).Size())     // 8
	fmt.Println(reflect.TypeOf("Hello").Size()) // 16

	// 通过Comparable方法可以判断一个类型是否可以被比较
	fmt.Println(reflect.TypeOf("hello world!").Comparable())
	fmt.Println(reflect.TypeOf(1024).Comparable())
	fmt.Println(reflect.TypeOf([]int{}).Comparable())
	fmt.Println(reflect.TypeOf(struct{}{}).Comparable())

	// 使用ConvertibleTo方法可以判断一个类型是否可以被转换为另一个指定的类型
	/*
		type MyInterface interface {
		  My() string
		}

		type MyStruct struct {
		}

		func (m MyStruct) My() string {
		  return "my"
		}

		type HisStruct struct {
		}

		func (h HisStruct) String() string {
		  return "his"
		}

		func main() {
		  rIface := reflect.TypeOf(new(MyInterface)).Elem()
		  fmt.Println(reflect.TypeOf(new(MyStruct)).Elem().ConvertibleTo(rIface)) // true
		  fmt.Println(reflect.TypeOf(new(HisStruct)).Elem().ConvertibleTo(rIface)) // false
		}
	*/

	// reflect.Value代表着反射接口的值，使用reflect.ValueOf()函数可以将变量转换成reflect.Value
	reflectValue := reflect.ValueOf(str)
	fmt.Println(reflectValue)

	// Type方法可以获取一个反射值的类型
	num := 114514
	rValue := reflect.ValueOf(num)
	fmt.Println(rValue, rValue.Type()) // int

	/*
		// 返回一个表示v地址的指针反射值
		func (v Value) Addr() Value

		// 返回一个指向v的原始值的uinptr 等价于 uintptr(Value.Addr().UnsafePointer())
		func (v Value) UnsafeAddr() uintptr

		// 返回一个指向v的原始值的uintptr
		// 仅当v的Kind为 Chan, Func, Map, Pointer, Slice, UnsafePointer时，否则会panic
		func (v Value) Pointer() uintptr

		// 返回一个指向v的原始值的unsafe.Pointer
		// 仅当v的Kind为 Chan, Func, Map, Pointer, Slice, UnsafePointer时，否则会panic
		func (v Value) UnsafePointer() unsafe.Pointer
	*/
	ele := reflect.ValueOf(&num).Elem()
	fmt.Println("&num", &num)
	fmt.Println("Addr", ele.Addr())
	fmt.Println("UnsafeAddr", unsafe.Pointer(ele.UnsafeAddr()))
	fmt.Println("Pointer", unsafe.Pointer(ele.Addr().Pointer()))
	fmt.Println("UnsafePointer", ele.Addr().UnsafePointer())

	// 倘若通过反射来修改反射值，那么其值必须是可取址的，这时应该通过指针来修改其元素值，而不是直接尝试修改元素的值
	number := new(int)
	*number = 114514
	rValue2 := reflect.ValueOf(number)
	// 获取指针指向的元素
	ele = rValue2.Elem()
	ele.SetInt(1919)
	// 通过Interface()方法可以获取反射值原有的值
	fmt.Println(ele, ele.Interface()) // fmt.Println会反射获取参数的类型，如果是reflect.Value类型的话，会自动调用Value.Interface()来获取其原始值
}
