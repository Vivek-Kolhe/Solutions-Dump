for _ in 0...(gets.chomp.to_i)
  n = gets.chomp.to_i
  arr = gets.chomp.split().map{|x| x.to_i}
  height, flag = 1, true
  for i in 1...(n + 1)
    if (arr[i] == 0) and (arr[i-1] == 0)
      puts -1
      flag = false
      break
    elsif (arr[i] == 1) and (arr[i-1] == 1)
      height += 5
    elsif arr[i-1] == 1
      height += 1
    end
  end
  if flag
    puts height
  end
end
