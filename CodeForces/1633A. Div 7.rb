n = gets.chomp.to_i
for _ in 0...n
  num = gets.chomp
  len = num.length()
  if num.to_i % 7 == 0
    puts num
  else
    for i in 1..9
      temp = (num[0, len-1] + i.to_s).to_i
      if temp % 7 == 0
        puts temp
        break
      end
    end
  end
end
