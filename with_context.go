// 终止正在运行的脚本
package main

import (
	"context"
	"time"

	lua "github.com/yuin/gopher-lua"
)

func stop_lua() {
	L := lua.NewState()
	defer L.Close()
	// 设置一秒后超时
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// set the context to our LState
	L.SetContext(ctx)
	// 执行lua 但是会被context中断
	L.DoFile("with_context.lua")
}
