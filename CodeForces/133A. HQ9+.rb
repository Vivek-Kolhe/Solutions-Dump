string = gets.chomp
instructions, flag = "HQ9", false
for i in 0...3
  if string.include?instructions[i]
    puts "YES"
    flag = true
    break
  end
end
if not flag
  puts "NO"
end
