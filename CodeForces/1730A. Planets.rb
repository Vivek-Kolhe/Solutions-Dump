for _ in 0...gets.chomp.to_i
  n, c = gets.chomp.split(' ').map{|x| x.to_i}
  arr = gets.chomp.split(' ').map{|x| x.to_i}
  uniq = arr.uniq
  cost = 0
  for i in 0...uniq.length
    temp = arr.count(uniq[i])
    if temp < c
      cost += temp
    else
      cost += c
    end
  end
  puts cost
end
