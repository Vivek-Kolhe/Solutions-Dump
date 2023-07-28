n, q = gets.chomp.split().map{ |x| x.to_i }
arr = gets.chomp.split().map{ |x| x.to_i }.sort
sums = [0]
for val in arr
  sums.append(sums[-1] + val)
end
for _ in 0...q
  x, y = gets.chomp.split().map{ |x| x.to_i }
  puts sums[n-x+y] - sums[n-x]
end
