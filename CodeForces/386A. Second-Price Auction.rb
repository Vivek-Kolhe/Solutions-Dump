n = gets.chomp.to_i
arr = gets.chomp.split().map{|x| x.to_i}
temp = arr.sort{|a, b| b <=> a}
print "#{arr.index(temp[0]) + 1} #{temp[1]}"
