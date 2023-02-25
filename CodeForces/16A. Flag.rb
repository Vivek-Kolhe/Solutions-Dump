flag = true
n, m = gets.chomp.split(' ').map{ |x| x.to_i}
flag_ = []
for _ in 0...n
  inp = gets.chomp
  if flag_[-1] == inp
    flag = false
  end
  if inp.count(inp[0]) < m
    flag = false
  end
  flag_.append(inp)
end
if not flag
  puts 'NO'
else
  puts 'YES'
end
