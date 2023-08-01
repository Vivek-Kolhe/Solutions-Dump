gets.chomp.to_i.times do
  n = gets.chomp.to_i
  dialogue = gets.chomp
  stack = []
  for i in 0...n
    if dialogue[i] == 'Q'
      stack.append(dialogue[i])
    else
      if stack.length > 0
        stack.pop()
      end
    end
  end
  if stack.length > 0
    puts 'No'
  else
    puts 'Yes'
  end
end
