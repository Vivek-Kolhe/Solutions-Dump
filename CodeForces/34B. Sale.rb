m, n = gets.chomp.split().map{|x| x.to_i}
arr = gets.chomp.split().map{|x| x.to_i}.sort
sum_ = 0
if arr[0] >= 0
  puts sum_
else
  for i in 0...n
    if arr[i] < 0
      sum_ += -(arr[i])
    else
      break
    end
  end
  puts sum_
end
