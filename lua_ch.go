// 在lua中使用ch
package main

import (
	"time"

	lua "github.com/yuin/gopher-lua"
)

// 在lua中使用ch的例子
func lua_ch() {
	L := lua.NewState()
	defer L.Close()

	ch1 := make(chan lua.LValue)
	ch2 := make(chan lua.LValue)
	L.SetGlobal("ch1", lua.LChannel(ch1))
	L.SetGlobal("ch2", lua.LChannel(ch2))

	// 通过ch1发送若干数据
	go func() {
		for x := 0; x < 10; x++ {
			ch1 <- lua.LNumber(x)
			time.Sleep(3 * time.Second)
		}
		close(ch1)
	}()

	// 通过ch2接收若干数据
	go func() {
		for x := 0; x < 10; x++ {
			<-ch2
			time.Sleep(3 * time.Second)
		}
		close(ch2)
	}()

	// 执行lua
	L.DoFile("lua_ch.lua")
}
