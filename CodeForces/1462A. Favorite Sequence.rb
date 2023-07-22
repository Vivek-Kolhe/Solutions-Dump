for _ in 0...gets.chomp.to_i
  n = gets.chomp.to_i
  num = gets.chomp.split(' ')
  res = []
  ind = 1
  while num.length > 0
    if ind.odd?
      res.append(num[0])
      num = num[1..]
    else
      res.append(num[-1])
      num = num[..-2]
    end
    ind += 1
  end
  puts res.join(' ')
end
