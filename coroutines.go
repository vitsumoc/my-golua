// 在go中使用协程
package main

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func coroutines() {
	L := lua.NewState()
	defer L.Close()

	// 定义coro函数
	L.DoString(`
    function coro()
      local i = 0
        while i < 100 do
          coroutine.yield(i)
          i = i+1
        end
      return i
    end
	`)

	co, _ := L.NewThread()                     /* create a new thread */
	fn := L.GetGlobal("coro").(*lua.LFunction) /* get function from lua */
	for {
		st, err, values := L.Resume(co, fn)
		if st == lua.ResumeError {
			fmt.Println("yield break(error)")
			fmt.Println(err.Error())
			break
		}

		for i, lv := range values {
			fmt.Printf("%v : %v\n", i, lv)
		}

		if st == lua.ResumeOK {
			fmt.Println("yield break(ok)")
			break
		}
	}
}
