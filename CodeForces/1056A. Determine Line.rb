n = gets.chomp.to_i
arr, res = [], []
for _ in 0...n
  arr.append(gets.chomp.split().map{|i| i.to_i})
end
arr.sort!{|a, b| a.length() <=> b.length()}
for i in 1...arr[0].length()
  temp_bool = []
  for j in 1...n
    if arr[j].include?arr[0][i]
      temp_bool.append(true)
    end
  end
  if temp_bool.count(true) == n - 1
    res.append(arr[0][i])
  end
end
puts res.uniq.join(" ")
