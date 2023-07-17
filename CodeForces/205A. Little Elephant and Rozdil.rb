n = gets.chomp.to_i
arr = gets.chomp.split(' ').map{ |x| x.to_i }
min = arr.min
if arr.count(min) > 1
  puts "Still Rozdil"
else
  puts arr.find_index(min) + 1
end
