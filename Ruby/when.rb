x = 200
y = true
c = 0
for i in 1 .. 25
  if y == true
    y = false
    c += 5
  else
    y = true
    c += 10
  end
end
puts c