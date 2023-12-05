/*
The LState is not goroutine-safe.
It is recommended to use one LState per goroutine and communicate between goroutines by using channels.

Channels are represented by channel objects in GopherLua. And a channel table provides functions for performing channel operations.

Some objects can not be sent over channels due to having non-goroutine-safe objects inside itself.

a thread(state)
a function
an userdata
a table with a metatable
You must not send these objects from Go APIs to channels.
*/
package main

import (
	"time"

	lua "github.com/yuin/gopher-lua"
)

func receiver(ch, quit chan lua.LValue) {
	L := lua.NewState()
	defer L.Close()
	L.SetGlobal("ch", lua.LChannel(ch))
	L.SetGlobal("quit", lua.LChannel(quit))
	if err := L.DoString(`
	local exit = false
	while not exit do
		channel.select(
			{"|<-", ch, function(ok, v)
				if not ok then
					print("channel closed")
					exit = true
				else
					print("received:", v)
				end
			end},
			{"|<-", quit, function(ok, v)
					print("quit")
					exit = true
			end}
		)
	end
`); err != nil {
		panic(err)
	}
}

func sender(ch, quit chan lua.LValue) {
	L := lua.NewState()
	defer L.Close()
	L.SetGlobal("ch", lua.LChannel(ch))
	L.SetGlobal("quit", lua.LChannel(quit))
	if err := L.DoString(`
	ch:send("1")
	ch:send("2")
`); err != nil {
		panic(err)
	}
	ch <- lua.LString("3")
	quit <- lua.LTrue
}

func goroutines() {
	ch := make(chan lua.LValue)
	quit := make(chan lua.LValue)
	go receiver(ch, quit)
	go sender(ch, quit)
	time.Sleep(3 * time.Second)
}
