YES = 'Yes' * 20
for _ in 0...gets.chomp.to_i
  string = gets.chomp
  if YES.include?string
    puts 'YES'
  else
    puts 'NO'
  end
end
