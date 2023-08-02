gets.chomp.to_i.times do
  n = gets.chomp.to_i
  word = gets.chomp
  i = 1
  flag = false
  while i < n
    if word[i] != word[i+1]
      flag = true
      break
    end
    i += 3
  end
  if flag
    puts 'NO'
  else
    puts 'YES'
  end
end
