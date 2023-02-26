# cook your code here
for _ in 0...gets.chomp.to_i
  a, b = gets.chomp.split(' ').map{|x| x.to_i}
  if (a + b > 6)
    puts 'YES'
  else
    puts 'NO'
  end
end
