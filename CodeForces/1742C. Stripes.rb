for _ in 0...gets.chomp.to_i
  board = []
  gets
  for i in 0...8
    board.append(gets.chomp)
  end
  if board.include?'R'*8
    puts 'R'
  else
    puts 'B'
  end
end
