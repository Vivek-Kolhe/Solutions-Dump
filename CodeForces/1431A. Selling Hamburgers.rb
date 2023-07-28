for _ in 0...gets.chomp.to_i
  n = gets.chomp.to_i
  coins = gets.chomp.split().map{ |x| x.to_i }.sort
  res = []
  for i in 0...n
    res.append((n - i) * coins[i])
  end
  puts res.max
end
