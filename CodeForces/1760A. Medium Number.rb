for _ in 0...gets.chomp.to_i
  puts gets.chomp.split(' ').map{|x| x.to_i}.sort[1]
end
