a = gets.chomp
if a == a.reverse
  puts a
else
  a += a.reverse
  puts a
end
