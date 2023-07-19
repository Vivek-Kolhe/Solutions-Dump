panels = ""
k = gets.chomp.to_i
for _ in 0...4
  panels += gets.chomp.gsub('.', '')
end
panels = panels.split('')
flag = false
uniq = panels.uniq
for i in 0...uniq.length
  if panels.count(uniq[i]) > (2 * k)
    flag = true
    break
  end
end
if flag
  puts "NO"
else
  puts "YES"
end
