original = gets.chomp
arr = []
for _ in 0...gets.chomp.to_i
  arr.append(gets.chomp)
end
if arr.include?original
  puts "YES"
else
  flag = false
  for i in arr
    if i.end_with?original[0]
      for j in arr
        if j.start_with?original[-1]
          flag = true
          break
        end
      end
    end
    if flag
      puts "YES"
      break
    end
  end
  if not flag
    puts "NO"
  end
end
