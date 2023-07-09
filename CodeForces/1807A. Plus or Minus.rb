for i in 0...gets.chomp.to_i
  a, b, c = gets.chomp.split(" ").map{ |x| x.to_i }
  if a + b == c
    puts '+'
  elsif a - b == c
    puts '-'
  end
end
