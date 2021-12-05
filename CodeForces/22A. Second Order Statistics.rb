n = gets.chomp.to_i
arr = gets.chomp.split().map{|x| x.to_i}.uniq.sort
if arr.length() == 1
  puts "NO"
else
  puts arr[1]
end
