
n = gets.chomp.to_i
arr = gets.chomp.split(" ").map{|item| item.to_i}
new_arr, count = arr.uniq, []
for i in new_arr
  count.append(arr.count(i))
end
puts "#{count.max} #{new_arr.length()}"
