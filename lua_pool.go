// lua虚拟机池
package main

import (
	"sync"

	lua "github.com/yuin/gopher-lua"
)

func lua_pool() {
	defer luaPool.Shutdown()
	// 在协程中使用lua环境
	go MyWorker()
	go MyWorker()
	/* etc... */
}

// 假设是调用lua虚拟机的方法
func MyWorker() {
	L := luaPool.Get()
	defer luaPool.Put(L)
	/* your code here */
}

// lua池数据结构
type lStatePool struct {
	m     sync.Mutex
	saved []*lua.LState
}

// 按顺序获得虚拟机, 不够时可以增加
func (pl *lStatePool) Get() *lua.LState {
	pl.m.Lock()
	defer pl.m.Unlock()
	n := len(pl.saved)
	if n == 0 {
		return pl.New()
	}
	x := pl.saved[n-1]
	pl.saved = pl.saved[0 : n-1]
	return x
}

// 虚拟机初始化方法, 按照业务可以设置初始化数据
func (pl *lStatePool) New() *lua.LState {
	L := lua.NewState()
	// setting the L up here.
	// load scripts, set global variables, share channels, etc...
	return L
}

// 归还虚拟机的方法
func (pl *lStatePool) Put(L *lua.LState) {
	pl.m.Lock()
	defer pl.m.Unlock()
	pl.saved = append(pl.saved, L)
}

// 关闭所有的虚拟机
func (pl *lStatePool) Shutdown() {
	for _, L := range pl.saved {
		L.Close()
	}
}

// 初始容量4
var luaPool = &lStatePool{
	saved: make([]*lua.LState, 0, 4),
}
