a = gets.chomp.to_i
b = gets.chomp.to_i
c = gets.chomp.to_i
while true
  if (b >= a * 2) and (c >= a * 4)
    b = a * 2
    c = a * 4
    puts a + b + c
    break
  else
    a -= 1
  end
end
