n = gets.chomp.to_i
arr = gets.chomp.split().map{|x| x.to_i}
a, b = gets.chomp.split().map{|x| x.to_i}
sum_ = 0
for i in (a - 1)...(b - 1)
  sum_ += arr[i]
end
puts sum_
