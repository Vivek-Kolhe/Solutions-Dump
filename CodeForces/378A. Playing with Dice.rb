a, b = gets.chomp.split().map{|x| x.to_i}
win_a, draw, win_b = 0, 0, 0
for i in 1...7
  if (i - a).abs < (i - b).abs
    win_a += 1
  elsif (i - a).abs > (i - b).abs
    win_b += 1
  else
    draw += 1
  end
end
puts "#{win_a} #{draw} #{win_b}"
