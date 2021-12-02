p, n = gets.chomp.split().map{|x| x.to_i}
table = {}
for i in 0...p
  table[i] = nil
end
flag, ind = true, -1
for _ in 0...n
  num = gets.chomp.to_i
  val = num % p
  if (table[val] != nil) and (flag)
    ind = _ + 1
    flag = false
  else
    table[val] = num
  end
end
if flag
  puts -1
else
  puts ind
end
