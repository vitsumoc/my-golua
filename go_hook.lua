-- print("i'm lua")
-- print(onStart)

local _myStart = function(dog)
  -- print(dog)
  print("lua:" .. dog:name())
end

local _myPrint = function(dog)
  dog:name("zhao cai")
end

onStart(_myStart)
onPrint(_myPrint)