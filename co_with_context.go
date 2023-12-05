// 使用context控制协程
package main

import (
	"context"
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func co_with_context() {
	L := lua.NewState()
	defer L.Close()
	ctx, cancel := context.WithCancel(context.Background())
	L.SetContext(ctx)
	defer cancel()
	L.DoFile("co_with_context.lua")

	// 新协程
	co, cocancel := L.NewThread()
	defer cocancel()
	fn := L.GetGlobal("coro").(*lua.LFunction)

	// 执行
	_, err, values := L.Resume(co, fn) // err is nil
	fmt.Println(err, values)

	cancel() // cancel the parent context

	// 协程也会被关闭
	_, err, values = L.Resume(co, fn)
	fmt.Println(err, values)
}
