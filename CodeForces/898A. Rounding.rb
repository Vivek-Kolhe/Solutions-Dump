n = gets.chomp
if n[-1] == "5"
  puts n.to_i - 5
elsif n.to_i % 10 == 0
  puts n
else
  if n[-1].to_i < 5
    puts n.to_i - n[-1].to_i
  elsif n[-1].to_i > 5
    puts n.to_i + (10 - n[-1].to_i)
  end
end
