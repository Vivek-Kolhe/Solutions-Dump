t = gets.chomp.to_i
for _ in 0...t
  string = gets.chomp
  count_a, count_b, count_c = string.count("A"), string.count("B"), string.count("C")
  if count_a + count_c == count_b
    puts "YES"
  else
    puts "NO"
  end
end
