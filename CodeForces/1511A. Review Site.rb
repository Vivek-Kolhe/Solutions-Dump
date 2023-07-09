for _ in 0...gets.chomp.to_i
  n = gets.chomp.to_i
  arr = gets.chomp.split(" ").map{ |x| x.to_i }
  ups, downs = arr.count(1), arr.count(2)
  mid = n - (ups + downs)
  puts ups + mid
end
