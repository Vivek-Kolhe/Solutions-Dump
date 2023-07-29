n = gets.chomp.to_i
res = []
i = 1
while i <= n
  temp = (n - i) / 2
  res.append('*' * temp + 'D' * i + '*' * temp)
  i += 2
end
res += res[0, res.length-1].reverse
puts res
