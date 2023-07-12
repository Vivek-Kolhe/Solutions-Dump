n = gets.chomp.to_i
arr = gets.chomp.split(' ').map{ |x| x.to_i }
count = 1
res = []
for i in 0...(n - 1)
  if (arr[i] <= arr[i+1])
    count += 1
  else
    res.append(count)
    count = 1
  end
end
res.append(count)
puts res.max
