for _ in 0...gets.chomp.to_i
  a, b, c = gets.chomp.split(' ').map{ |x| x.to_i}
  if (a + b == c) or (b + c == a) or (c + a == b)
    puts 'YES'
  else
    puts 'NO'
  end
end
