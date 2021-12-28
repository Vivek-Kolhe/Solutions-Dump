n = gets.chomp.to_i
flag = false
for i in 1...(n + 1)
  for j in 1...(n + 1)
    if (n < i * j) and (i % j == 0)
      puts "#{i} #{j}"
      flag = true
      break
    end
  end
  if flag
    break
  end
end
if not flag
  puts -1
end
