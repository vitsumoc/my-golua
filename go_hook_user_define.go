// go钩子示例中的用户定义数据内容
package main

import lua "github.com/yuin/gopher-lua"

// 定义普通的golang结构, 对应golang业务
type Dog struct {
	Name string
}

// 为用户结构命名
const luaDogTypeName = "Dog"

// 将用户结构注册到Lua的方法
func registerDogType(L *lua.LState) {
	// 元表
	mt := L.NewTypeMetatable(luaDogTypeName)
	// 把元表注册上去
	L.SetGlobal("Dog", mt)
	// 使用元表, 将Dog类的各项查询都指向golang
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), dogMethods))
}

// 校验dog类型
func checkDog(L *lua.LState) *Dog {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Dog); ok {
		return v
	}
	L.ArgError(1, "dog expected")
	return nil
}

// 对name字段的存取方法
func dogName(L *lua.LState) int {
	d := checkDog(L)
	// 这是set 会有两个参数 ud 和 参数name
	if L.GetTop() == 2 {
		d.Name = L.CheckString(2)
		return 0
	}
	// 这是get 只有一个参数 就是ud
	L.Push(lua.LString(d.Name))
	return 1
}

// 描述dog可用方法的列表 属性名<->存取方法
var dogMethods = map[string]lua.LGFunction{
	"name": dogName,
}
