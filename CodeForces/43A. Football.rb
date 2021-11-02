arr = []
for _ in 0...gets.chomp.to_i
  arr.append(gets.chomp)
end
teams = arr.uniq
if arr.count(teams[0]) > arr.count(teams[-1])
  puts teams[0]
else
  puts teams[-1]
end
