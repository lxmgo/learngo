package main

import (
	"reflect"
	"fmt"
)

type ControllerInterface interface {
	Init(action string, method string)
}

type Controller struct {
	Action string
	Method string
	Tag string `json:"tag"`
}

func (c *Controller) Init(action string, method string){
	c.Action = action
	c.Method = method

	fmt.Println("Init() is run.")
	fmt.Println("c:",c)
}

func (c *Controller) Test(){
	fmt.Println("Test() is run.")
}


func main(){
	//初始化
	runController := &Controller{
		Action:"Run1",
		Method:"GET",
	}

	//Controller实现了ControllerInterface方法,因此它就实现了ControllerInterface接口
	var i ControllerInterface
	i = runController

	// 得到实际的值,通过v我们获取存储在里面的值,还可以去改变值
	v := reflect.ValueOf(i)
	fmt.Println("value:",v)

	// 得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	t := reflect.TypeOf(i)
	fmt.Println("type:",t)

	// 转化为reflect对象之后我们就可以进行一些操作了,也就是将reflect对象转化成相应的值,例如
	controllerType := t.Elem()
	tag := controllerType.Field(2).Tag
	fmt.Println("Tag:", tag)

	// 获取i所指向的对象的类型(reflect.Value)
	controllerValue := v.Elem()
	fmt.Println("controllerType(reflect.Value):",controllerType)
	//获取存储在第一个字段里面的值
	fmt.Println("Action:", controllerValue.Field(0).String())

	method, _ := t.MethodByName("Init")
	fmt.Println(method)

	vMethod := v.MethodByName("Init")
	fmt.Println(vMethod)

	// 有输入参数的方法调用
	// 构造输入参数
	args1 := []reflect.Value{reflect.ValueOf("Run2"),reflect.ValueOf("POST")}
	// 通过v进行调用
	v.MethodByName("Init").Call(args1)

	// 无输入参数的方法调用
	// 构造zero value
	args2 := make([]reflect.Value, 0)
	// 通过v进行调用
	v.MethodByName("Test").Call(args2)

}
