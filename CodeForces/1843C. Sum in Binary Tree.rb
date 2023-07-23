for _ in 0...gets.chomp.to_i
  n = gets.chomp.to_i
  sum = n
  while n > 0
    temp = n / 2
    sum += temp
    if n.odd?
      n -= temp + 1
    else
      n -= temp
    end
  end
  puts sum
end
