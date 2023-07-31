gets.chomp.to_i.times do
  n = gets.chomp.to_i
  arr = gets.chomp.split().map(&:to_i)
  res = []
  max1, max2 = arr.max(2)
  for i in 0...n
    if arr[i] == max1
      res.append(arr[i] - max2)
    else
      res.append(arr[i] - max1)
    end
  end
  puts res.join(' ')
end
