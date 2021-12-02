n, d = gets.chomp.split().map{|x| x.to_i}
arr = gets.chomp.split().map{|x| x.to_i}.sort
count = 0
for i in 0...n
  for j in (i + 1)...n
    diff = arr[j] - arr[i]
    if (diff >= 0) and (diff <= d)
      count += 1
    end
  end
end
puts count * 2
