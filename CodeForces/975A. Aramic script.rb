n = gets.chomp.to_i
arr = gets.chomp.split(" ")
result = []
for i in 0...n
  temp = arr[i].split(//).uniq.sort.join("")
  if not result.include?temp
    result.append(temp)
  end
end
puts result.length()
