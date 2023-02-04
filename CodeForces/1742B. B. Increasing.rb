for _ in 0...gets.chomp.to_i
  n = gets.chomp.to_i
  arr = gets.chomp.split(' ').map{|x| x.to_i}
  flag = true
  for i in 0...n
    if arr.count(arr[i]) > 1
      flag = false
      puts 'NO'
      break
    end
  end
  if flag
    puts 'YES'
  end
end
