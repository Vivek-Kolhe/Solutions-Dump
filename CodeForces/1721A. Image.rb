for _ in 0...gets.chomp.to_i
  matrix = []
  for i in 0...2
    temp = gets.chomp.split('')
    for color in temp
      matrix.append(color)
    end
  end
  matrix.uniq!
  puts matrix.length() - 1
end
