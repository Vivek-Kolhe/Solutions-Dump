n = gets.chomp.to_i
size = gets.chomp.to_i
arr = []
for _ in 0..n
  arr.append(gets.to_i)
end
arr, sum_ = arr.sort{|a, b| b <=> a}, 0
for i in 0..n
  sum_ += arr[i]
  if sum_ >= size
    puts i + 1
    break
  end
end
