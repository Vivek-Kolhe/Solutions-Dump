n = gets.chomp.to_i
if n <= 10 or n >= 22
  puts 0
else
  if n == 21
    puts 4
  elsif n == 20
    puts 15
  else
    puts 4
  end
end
