n = gets.chomp.to_i
arr = gets.chomp.split(' ').map{ |x| x.to_i }
arr.uniq!
puts arr.length
