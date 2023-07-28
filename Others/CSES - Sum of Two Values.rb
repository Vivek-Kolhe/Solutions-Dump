n, x = gets.chomp.split().map{ |x| x.to_i }
arr = gets.chomp.split().map{ |x| x.to_i }
map = {}
flag = false
for i in 0...n
  comp = x - arr[i]
  if map.has_key?comp
    puts [map[comp] + 1, i + 1].join(" ")
    flag = true
    break
  else
    map[arr[i]] = i
  end
end
if not flag
  puts 'IMPOSSIBLE'
end
