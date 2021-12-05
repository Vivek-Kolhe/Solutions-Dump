n = gets.chomp.to_i
arr = gets.chomp.split().map{|x| x.to_i}
left, right = arr.count(0), arr.count(1)
for i in 0...n
  if arr[i] == 0
    left -= 1
  else
    right -= 1
  end
  if (left == 0) or (right == 0)
    puts i + 1
    break
  end
end
