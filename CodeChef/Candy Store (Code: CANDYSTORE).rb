# cook your code here
for _ in 0...gets.chomp.to_i
  x, y = gets.chomp.split(' ').map{|x| x.to_i}
  if y <= x
    puts y
  else
    puts ((x * 1) + ((y - x) * 2))
  end
end
