n = gets.chomp.to_i
arr = gets.chomp.split(' ').map{|x| x.to_f}.sort
average = arr.sum / n
count, ind = 0, 0
while average < 4.5
  arr[ind] = 5.to_f
  ind += 1
  count += 1
  average = arr.sum / n
end
puts count
