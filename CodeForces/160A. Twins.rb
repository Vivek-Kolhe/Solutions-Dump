n = gets.chomp.to_i
arr = gets.chomp.split(' ').map{ |x| x.to_i }.sort{ |a, b| b <=> a }
stole = 0
for i in 0...n
  if arr[0, i+1].sum > arr[i+1, n].sum
    stole += 1
    break
  else
    stole += 1
  end
end
puts stole
