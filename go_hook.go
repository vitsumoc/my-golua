// 在golang中提供钩子, 使lua可以注册脚本, 在脚本中获得并修改用户数据
package main

import (
	"fmt"
	"time"

	lua "github.com/yuin/gopher-lua"
)

func go_hook() {
	L := lua.NewState()
	defer L.Close()
	// 把狗狗注册给lua
	registerDogType(L)
	// 把钩子提供给lua
	L.SetGlobal("onStart", L.NewFunction(regStart))
	L.SetGlobal("onPrint", L.NewFunction(regPrint))
	// 执行lua文件
	err := L.DoFile("go_hook.lua")
	if err != nil {
		panic(err)
	}

	// 模拟业务循环
	// 被测试使用的狗狗
	var dog Dog = Dog{
		Name: "pika",
	}
	// 用userData包裹
	ud := L.NewUserData()
	ud.Value = &dog
	L.SetMetatable(ud, L.GetTypeMetatable(luaDogTypeName))

	for {
		time.Sleep(2 * time.Second)
		fmt.Println("golang 1:" + dog.Name)

		// 执行第一个钩子
		for x := 0; x < len(startFuncs); x++ {
			err := L.CallByParam(lua.P{
				Fn:      startFuncs[x],
				NRet:    0, // 返回参数个数
				Protect: true,
			}, ud)
			if err != nil {
				panic(err)
			}
		}
		fmt.Println("golang 2:" + dog.Name)

		// 执行第二个钩子
		for x := 0; x < len(printFuncs); x++ {
			err := L.CallByParam(lua.P{
				Fn:      printFuncs[x],
				NRet:    0, // 返回参数个数
				Protect: true,
			}, ud)
			if err != nil {
				panic(err)
			}
		}
		fmt.Println("golang 3:" + dog.Name)
	}
}

// 本示例使用两个钩子, 一个位于循环开始时, 另一个位于两次打印中间
// 钩子提供Dog实例, 包含如下函数
// func name(string1) stirng2
// 其中string1非空时可以用来设置name, string2 会返回最新的name

// lua中的函数列表, 保存所有被注册的钩子函数
var startFuncs []*lua.LFunction = make([]*lua.LFunction, 0)
var printFuncs []*lua.LFunction = make([]*lua.LFunction, 0)

// 应提供一个golang的函数, 这个函数可以把lua的函数注册进去
// 被注册的函数会在golang程序运行的某个节点被按序
// 入参是一个lua函数
// 每一次注册都会增加一个执行函数
func regStart(L *lua.LState) int {
	luaFunc := L.ToFunction(1)
	startFuncs = append(startFuncs, luaFunc)
	return 1
}

func regPrint(L *lua.LState) int {
	luaFunc := L.ToFunction(1)
	printFuncs = append(printFuncs, luaFunc)
	return 1
}
