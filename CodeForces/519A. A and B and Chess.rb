wts = {
  'Q' => 9,
  'R' => 5,
  'B' => 3,
  'N' => 3,
  'P' => 1,
  'K' => 0
}
black, white = 0, 0
for _ in 0...8
  row = gets.chomp
  for i in 0...8
    char = row[i]
    if wts[char.upcase]
      if char == char.upcase
        white += wts[char]
      else
        black += wts[char.upcase]
      end
    end
  end
end
if black > white
  puts "Black"
elsif black < white
  puts "White"
else
  puts "Draw"
end
