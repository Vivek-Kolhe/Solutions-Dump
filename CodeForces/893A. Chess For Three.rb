def diff(a, b)
  for i in 0...b.length
    if not a.include?(b[i])
      return b[i]
    end
  end
end

n = gets.chomp.to_i
flag = true
wins = []
for _ in 0...n
  wins.append(gets.chomp)
end
game = '12'
for i in 0...n
  if game.include?(wins[i])
    game.sub!(diff(wins[i], game), diff(game, '123'))
  else
    flag = false
  end
end
if not flag
  puts 'NO'
else
  puts 'YES'
end
