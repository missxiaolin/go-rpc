package services

import (
	"reflect"
	"fmt"
	"services/test"
)

// 定义控制器函数Map类型，便于后续快捷使用
type ServiceMappersType map[string]map[string]reflect.Value

var ServiceMappers = make(ServiceMappersType, 0)

func registerService(name string, service ServiceInterface) {

	vf := reflect.ValueOf(service)
	vft := vf.Type()
	mNum := vf.NumMethod()
	fmt.Println("NumMethod:", mNum)

	// 遍历路由器的方法，并将其存入控制器映射变量中
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		fmt.Println("index:", i, " MethodName:", mName)
		serviceByName, ok := ServiceMappers[name]
		if !ok {
			serviceByName = make(map[string]reflect.Value)
			ServiceMappers[name] = serviceByName
		}
		serviceByName[mName] = vf.Method(i)
	}
}

func Init() {
	registerService("test", &test.Test{})
}
