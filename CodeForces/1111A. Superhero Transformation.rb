s = gets.chomp
t = gets.chomp
vowels = ["a", "e", "i", "o", "u"]
if s.length() != t.length()
  puts "No"
else
  flags = []
  for i in 0...s.length()
    if (vowels.include?s[i] and vowels.include?t[i]) or (not vowels.include?s[i] and not vowels.include?t[i])
      flags.append(true)
    end
  end
  if flags.count(true) == s.length()
    puts "Yes"
  else
    puts "No"
  end
end
