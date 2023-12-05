// 用户自定义数据
package main

import (
	lua "github.com/yuin/gopher-lua"
)

// 普通的golang结构
type Person struct {
	Name string
}

// 为用户结构命名
const luaPersonTypeName = "person"

// 将用户结构注册到Lua的方法
func registerPersonType(L *lua.LState) {
	// 元表
	mt := L.NewTypeMetatable(luaPersonTypeName)
	// 把元表注册上去
	L.SetGlobal("person", mt)
	// 指定new方法，指向go函数
	L.SetField(mt, "new", L.NewFunction(newPerson))
	// 把person类的方法都指向go
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), personMethods))
}

// Person构造器
func newPerson(L *lua.LState) int {
	// 拿一个String参数作为person名
	person := &Person{L.CheckString(1)}
	// 用ud包裹person
	ud := L.NewUserData()
	ud.Value = person
	// 绑定元表, 以设置各类方法
	L.SetMetatable(ud, L.GetTypeMetatable(luaPersonTypeName))
	// 扔进去完事
	L.Push(ud)
	return 1
}

// 校验第一个参数是否是UserData, 其中内容是否是Person
// 如果是则返回 *Person
func checkPerson(L *lua.LState) *Person {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Person); ok {
		return v
	}
	L.ArgError(1, "person expected")
	return nil
}

// 对name字段的存取方法
func personGetSetName(L *lua.LState) int {
	p := checkPerson(L)
	// 这是set 会有两个参数 ud 和 参数name
	if L.GetTop() == 2 {
		p.Name = L.CheckString(2)
		return 0
	}
	// 这是get 只有一个参数 就是ud
	L.Push(lua.LString(p.Name))
	return 1
}

// 描述person可用方法的列表 属性名<->存取方法
var personMethods = map[string]lua.LGFunction{
	"name": personGetSetName,
}

func user_defined_type() {
	L := lua.NewState()
	defer L.Close()
	registerPersonType(L)
	L.DoFile("user_defined_type.lua")
}
