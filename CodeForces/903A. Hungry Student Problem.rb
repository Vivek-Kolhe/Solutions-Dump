temp = [1, 2, 4, 5, 8, 11]
for _ in 0...(gets.chomp.to_i)
  n = gets.chomp.to_i
  if temp.include?n
    puts "NO"
  else
    puts "YES"
  end
end
