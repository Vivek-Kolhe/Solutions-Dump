for _ in 0...gets.chomp.to_i
  n = gets.chomp.to_i
  moves = gets.chomp
  candy = [1, 1]
  ini = [0, 0]
  flag = false
  for i in 0...n
    if moves[i] == 'U'
      ini[1] += 1
    elsif moves[i] == 'D'
      ini[1] -= 1
    elsif moves[i] == 'R'
      ini[0] += 1
    else
      ini[0] -= 1
    end
    if candy == ini
      flag = true
      puts 'YES'
      break
    end
  end
  if not flag
    puts 'NO'
  end
end
