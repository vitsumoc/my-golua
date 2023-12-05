// 数据类型
package main

// LChannel 的出现真让人激动
// Type name	Go type	Type() value	Constants
// LNilType	(constants)	LTNil	LNil
// LBool	(constants)	LTBool	LTrue, LFalse
// LNumber	float64	LTNumber	-
// LString	string	LTString	-
// LFunction	struct pointer	LTFunction	-
// LUserData	struct pointer	LTUserData	-
// LState	struct pointer	LTThread	-
// LTable	struct pointer	LTTable	-
// LChannel	chan LValue	LTChannel	-

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func data_model() {
	L := lua.NewState()
	defer L.Close()

	L.Push(lua.LString("i am string"))
	// 获得栈顶数据 类型是LValue
	lv := L.Get(-1)
	// 使用LString进行类型判断
	if str, ok := lv.(lua.LString); ok {
		// LString 可以直接转 string
		fmt.Println(string(str))
	}
	// 也可以用Type函数进行类型判断
	if lv.Type() != lua.LTString {
		panic("string required.")
	}

	// 取出以后还能再获得
	lv = L.Get(-1)
	if str, ok := lv.(lua.LString); ok {
		fmt.Println(string(str))
	}

	// 弹出以后则不可获得
	L.Pop(1)
	lv = L.Get(-1)
	if str, ok := lv.(lua.LString); ok {
		fmt.Println(string(str))
	}

	// 可以放入表格
	table := L.NewTable()
	table.RawSet(lua.LNumber(1), lua.LNumber(1))
	L.Push(table)
	// 可以取出表格
	lv = L.Get(-1)
	// 可以判断是否是Table
	if tbl, ok := lv.(*lua.LTable); ok {
		fmt.Println(L.ObjLen(tbl))
	}
	L.Pop(1)

	// BOOL值的存取方法
	_bool := lua.LTrue
	L.Push(_bool)
	lv = L.Get(-1)
	if lv.Type() == lua.LTBool {
		fmt.Println(lv)
	}
	L.Pop(1)

	//
}
