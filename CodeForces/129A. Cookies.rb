n = gets.chomp.to_i
arr = gets.chomp.split(' ').map{ |x| x.to_i }
total = arr.sum()
count = 0
for i in 0...n
  rem = total - arr[i]
  if rem.even?
    count += 1
  end
end
puts count
