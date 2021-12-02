n = gets.chomp.to_i
arr = gets.chomp.split().map{|x| x.to_i}
count = 0
for i in 1...(n - 1)
  if (arr[i] > arr[i + 1] and arr[i] > arr[i - 1]) or (arr[i] < arr[i + 1] and arr[i] < arr[i - 1])
    count += 1
  end
end
puts count
