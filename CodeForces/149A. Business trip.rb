k = gets.chomp.to_i
arr = gets.chomp.split().map{|x| x.to_i}
arr = arr.sort{|a, b| b <=> a}
if k == 0
  puts 0
else
  count = 0
  for i in 0...12
    count += arr[i]
    if count >= k
      puts i + 1
      break
    end
  end
  if count < k
    puts -1
  end
end
