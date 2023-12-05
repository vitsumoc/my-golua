package main

import (
	"vc/go-lua/mymodule"

	lua "github.com/yuin/gopher-lua"
)

func lua_call_go_lib() {
	L := lua.NewState()
	defer L.Close()
	L.PreloadModule("mymodule", mymodule.Loader)
	if err := L.DoFile("lua_call_go_lib.lua"); err != nil {
		panic(err)
	}
}
