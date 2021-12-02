vowels = ["a", "e", "i", "o", "u"]
syl, poem = [5, 7, 5], []
for i in 0...3
  poem.append(gets.chomp)
end
flag = true
for i in 0...3
  count = 0
  for j in 0...poem[i].length()
    if vowels.include?poem[i][j]
      count += 1
    end
  end
  if count != syl[i]
    puts "NO"
    flag = false
    break
  end
end
if flag
  puts "YES"
end
