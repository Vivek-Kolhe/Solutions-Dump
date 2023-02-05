n = gets.chomp.to_i
arr = gets.chomp.split(' ').map{|x| x.to_i}
max, min = arr[0], arr[0]
count = 0
if arr.length() == 1
  puts 0
else
  for i in 0...n
    if arr[i] > max
      max = arr[i]
      count += 1
    elsif min > arr[i]
      min = arr[i]
      count += 1
    end
  end
  puts count
end
