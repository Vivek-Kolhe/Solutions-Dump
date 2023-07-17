def min(*nums)
  nums.min
end

n = gets.chomp.to_i
arr = gets.chomp.split(' ').map{ |x| x.to_i }.sort
puts min((arr[1] - arr[-1]).abs, (arr[0] - arr[-2]).abs)
