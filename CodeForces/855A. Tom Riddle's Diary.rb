arr = []
for i in 0...gets.chomp.to_i
  name = gets.chomp
  if arr.include?name and arr.find_index(name) < i
    puts "YES"
  else
    arr.append(name)
    puts "NO"
  end
end
