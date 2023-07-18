leaves = []
for _ in 0...gets.chomp.to_i
  leaf = gets.chomp
  if not leaves.include?(leaf)
    leaves.append(leaf)
  end
end
puts leaves.length
