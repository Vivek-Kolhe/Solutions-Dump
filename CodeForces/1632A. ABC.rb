for _ in 0...(gets.chomp.to_i)
  len = gets.chomp.to_i
  string = gets.chomp
  if len >= 3
    puts "NO"
  elsif (string == "00") or (string == "11")
    puts "NO"
  else
    puts "YES"
  end
end
