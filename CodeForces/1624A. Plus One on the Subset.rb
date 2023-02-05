for _ in 0...gets.chomp.to_i
  n = gets.chomp.to_i
  arr = gets.chomp.split(' ').map{|x| x.to_i}
  puts arr.max() - arr.min()
end
