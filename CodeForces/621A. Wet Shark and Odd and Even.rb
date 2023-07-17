n = gets.chomp.to_i
arr = gets.chomp.split(' ').map{ |x| x.to_i }.sort
sum = arr.sum
if sum % 2 == 0
  puts sum
else
  for i in 0...n
    if arr[i].odd?
      sum -= arr[i]
      break
    end
  end
  puts sum
end
