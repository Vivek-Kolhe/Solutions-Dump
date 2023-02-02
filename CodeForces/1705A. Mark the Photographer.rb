for _ in 0...gets.chomp.to_i
  n, x = gets.chomp.split(" ").map{ |x| x.to_i}
  arr = gets.chomp.split(" ").map{ |x| x.to_i}
  arr = arr.sort
  flag = false
  for i in 0...n
    if arr[i+n] - arr[i] < x
      flag = true
      break
    end
  end
  if flag
    puts 'NO'
  else
    puts 'YES'
  end
end
