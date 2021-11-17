for i in 0...gets.chomp.to_i
  n = gets.chomp.to_i
  arr = gets.chomp.split(" ").map{|item| item.to_i}
  temp = arr.uniq
  for i in temp
    if arr.count(i) == 1
      puts arr.index(i) + 1
    end
  end
end
