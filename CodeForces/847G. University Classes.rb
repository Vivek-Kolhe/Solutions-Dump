grps = []
for _ in 0...gets.chomp.to_i
  grps.append(gets.chomp)
end
count = 0
for i in 0...7
  temp = []
  for j in 0...grps.length
    temp.append(grps[j][i].to_i)
  end
  if count < temp.sum
    count = temp.sum
  end
end
puts count
