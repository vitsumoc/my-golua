-- 这是用select的方法
while true do
  local idx, recv, ok = channel.select(
    {"|<-", ch1, function(ok, data)
      print(ok, data)
    end},
    {"<-|", ch2, "ch2 lua value", function(data)
      print(data)
    end},
    {"default", function()
      -- print("default action")
    end}
  )

  if not ok then
    -- print("closed")
  elseif idx == 1 then -- received from ch1
    print(recv)
  elseif idx == 2 then -- received from ch2
    print(recv)
  end
end

-- 也可以用一些更加基本的函数 make send receive close
-- https://github.com/yuin/gopher-lua#lua-api