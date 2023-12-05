function coro()
  local i = 0
  while true do
    coroutine.yield(i)
        i = i+1
  end
  return i
end