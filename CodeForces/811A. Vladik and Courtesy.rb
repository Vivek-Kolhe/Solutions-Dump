a, b = gets.chomp.split().map{|x| x.to_i}
i = 1
while true
  if i % 2 != 0
    if a < i
      puts "Vladik"
      break
    end
    a -= i
  else
    if b < i
      puts "Valera"
      break
    end
    b -= i
  end
  i += 1
end
