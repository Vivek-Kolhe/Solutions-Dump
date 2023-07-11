def max(*nums)
  nums.max
end

n, c = gets.chomp.split(' ').map{ |x| x.to_i }
p = gets.chomp.split(' ').map{ |x| x.to_i }
t = gets.chomp.split(' ').map{ |x| x.to_i }
limak, radewoosh = 0, 0
lt, rt = 0, 0
for i in 1...(n+1)
  lt += t[i-1]
  limak += max(0, p[i-1] - (c * lt))
  rt += t[-i]
  radewoosh += max(0, p[-i] - (c * rt))
end
if limak > radewoosh
  puts 'Limak'
elsif limak < radewoosh
  puts 'Radewoosh'
else
  puts 'Tie'
end
