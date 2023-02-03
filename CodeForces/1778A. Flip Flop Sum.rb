for _ in 0...gets.chomp.to_i
  n = gets.chomp.to_i
  arr = gets.chomp.split(' ').map{ |x| x.to_i}
  flag = true
  if not arr.include?(-1)
    puts arr.sum - 4
  else
    for i in 0...n-1
      if (arr[i] == -1) and (arr[i+1] == -1)
        flag = false
        puts arr.sum + 4
        break
      end
    end
    if flag
      puts arr.sum
    end
  end
end
