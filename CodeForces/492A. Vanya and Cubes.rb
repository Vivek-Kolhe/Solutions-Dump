n = gets.chomp.to_i
height, count = 0, 0
while (count <= n)
  height += 1
  count += (height * (height - 1)) / 2
end
puts height - 2
