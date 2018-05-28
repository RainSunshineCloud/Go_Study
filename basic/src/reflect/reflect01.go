package main

/*
总结：在reflect中：
field代表值，
name表示变量名，
Index表示索引

两个包文件value.go 和 type.go 分别代表值和类型的处理方法
*/
import (
	"fmt"
	"reflect"
)

type s struct {
	ints   int
	str    string
	slices []interface{}
	arr    [5]interface{}
	ss     *int
	float  float32
	b      bool
}

func main() {
	s1 := s{ints: 1}
	p := 12
	s1.ss = &p
	s1.valuesField()
	s1.types()
	s1.valueMethod()
	s1.typesMethod()
}

//value包对应的方法
func (s1 *s) valuesField() {
	values := reflect.ValueOf(*s1) //不能传指针
	fmt.Println("value:", values)
	name1 := "ints"
	v1 := values.FieldByName(name1) //获取值
	v2 := values.Field(0)           //获取值
	v3 := values.Field(4).Elem()    //通过指针获取value(如果是指针）
	name2 := []int{0}
	v4 := values.FieldByIndex(name2) //没什么毛线用，跟field一个德行除了类型不一样
	v5 := values.Field(6).Bool()     //获取布尔值，没什么毛线用
	v6 := values.Field(3).Cap()      //获取容量值
	v7 := values.Kind()              //类型
	v8 := values.NumField()          //字段个数

	fmt.Println("v1:", v1, "v2:", v2, "v3:", v3, "v4", v4, "v5:", v5, "v6:", v6, "v7:", v7, "v8:", v8)
}

//type包对应的方法
func (s *s) types() {
	types := reflect.TypeOf(*s)
	fmt.Println("type", types)

	name1 := "ints"
	t1, ok := types.FieldByName(name1)

	t2 := types.Field(0)
	t8 := types.NumField()
	fmt.Println("t1:", t1, ok, "t2:", t2, "t8:", t8)
}

//
func (s1 *s) valueMethod() {
	v := reflect.ValueOf(s1)

	mv1 := v.NumMethod() //方法数量
	name2 := "Echo"
	mv2 := v.MethodByName(name2) //通过值获取方法
	parm := make([]reflect.Value, 1)
	parm[0] = reflect.ValueOf("100")
	mv3 := mv2.Call(parm)
	fmt.Println("mv1:", mv1, "mv2:", mv2, "mv3:", mv3)
}

func (s1 *s) typesMethod() {
	t := reflect.TypeOf(s1)
	fmt.Println("mt:", t)
	name2 := "Echo"
	mt2, ok := t.MethodByName(name2)
	fmt.Println("mt2:", mt2, ok, "in:")

}

func (s *s) Echo(str string) {
	fmt.Println(str)
}
