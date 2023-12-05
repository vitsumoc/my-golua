package main

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func stackInfo(ls *lua.LState) {
	fmt.Println("当前堆栈: " + ls.String())
}
