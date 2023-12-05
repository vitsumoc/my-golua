package main

func main() {
	// hello() // 最基础的用法
	// data_model() // 基础数据类型
	// lua_call_go() // 在lua中调用go方法
	// coroutines() // 在go中使用lua协程
	// sub_lib() // 示范如何手动开启模块
	// lua_call_go_lib() // 在lua中使用go模块
	// go_call_lua() // 在golang中调用lua方法
	user_defined_type() // 在lua中使用golang数据
	// stop_lua() // 通过context控制停止
	// co_with_context() // 在有协程的情况下使用context控制
	// sharing_lua_byte() // 共享lua文件字节码, 减少开销
	// goroutines() // 通过go协程跑lua的示例 可以把ch带到lua中 和相关限制
	// lua_ch() // 在lua中使用ch的例子
	// lua_pool() // lua虚拟机池
	// go_hook() // 在golang中提供钩子, 使lua可以注册脚本, 在脚本中获得并修改用户数据
}
