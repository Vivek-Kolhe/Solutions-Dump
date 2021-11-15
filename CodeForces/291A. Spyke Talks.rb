n = gets.chomp.to_i
arr = gets.chomp.split(" ")
arr = arr.map{|item| item.to_i}
new_arr = arr.uniq
count, flag = 0, true
for i in new_arr
  if i != 0
    if arr.count(i) == 2
      count += 1
    elsif arr.count(i) > 2
      flag = false
      break
    end
  end
end
if not flag
  puts -1
else
  puts count
end
