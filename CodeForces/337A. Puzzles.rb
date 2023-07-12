def min(*nums)
  nums.min
end

n, m = gets.chomp.split(' ').map{ |x| x.to_i }
arr = gets.chomp.split(' ').map{ |x| x.to_i }.sort
best = 10000000000000
for i in 0...(m - n + 1)
  best = min(best, arr[i+n-1] - arr[i])
end
puts best
