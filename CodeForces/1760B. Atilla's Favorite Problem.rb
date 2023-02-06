alphabets = ('a'..'z').to_a
for _ in 0...gets.chomp.to_i
  n = gets.chomp.to_i
  string = gets.chomp.split('').sort
  puts alphabets.find_index(string[-1]) + 1
end
