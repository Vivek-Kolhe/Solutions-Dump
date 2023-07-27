for _ in 0...gets.chomp.to_i
  n, k = gets.chomp.split(' ').map{ |x| x.to_i}
  count = 0
  while n > 0
    if n % k == 0
      n /= k
      count += 1
    else
      rem = n % k
      n -= rem
      count += rem
    end
  end
  puts count
end
