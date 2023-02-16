for _ in 0...gets.chomp.to_i
  n = gets.chomp.to_i
  string = gets.chomp.split('')
  ques = ('A'..'Z').to_a
  balloons = 0
  for letter in string
    if ques.include?letter
      balloons += 2
      ques.delete(letter)
    else
      balloons += 1
    end
  end
  puts balloons
end
