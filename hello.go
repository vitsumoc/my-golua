// 最基础的用法
package main

import lua "github.com/yuin/gopher-lua"

func hello() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(`print("hello")`); err != nil {
		panic(err)
	}
}
