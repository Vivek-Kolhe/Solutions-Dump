n, k = gets.chomp.split().map{|x| x.to_i}
arr = gets.chomp.split().map{|x| x.to_i}
min_ = 99999
for i in 0...n
  temp = k % arr[i]
  if temp == 0
    if k / arr[i] < min_
      min_ = k / arr[i]
    end
  end
end
puts min_
