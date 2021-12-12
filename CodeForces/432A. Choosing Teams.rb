n, k = gets.chomp.split().map{|x| x.to_i}
arr = gets.chomp.split().map{|x| x.to_i}.sort
count_a, count_b = 0, 0
for i in 0...n
  if arr[i] <= 5 - k
    count_a += 1
  end
  if count_a == 3
    count_a = 0
    count_b += 1
  end
end
puts count_b
