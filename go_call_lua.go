// 在golang中使用lua
package main

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func go_call_lua() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile("go_call_lua.lua"); err != nil {
		panic(err)
	}
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("double"),
		NRet:    1, // 返回参数个数
		Protect: true,
	}, lua.LNumber(10)); err != nil {
		panic(err)
	}
	// 收到的结果
	ret := L.Get(-1)
	// 清理栈
	L.Pop(1)

	fmt.Println(ret)
}
