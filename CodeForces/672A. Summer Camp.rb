n = gets.chomp.to_i
string, num = "", "1"
while true
  string += num
  num = (num.to_i + 1).to_s
  if string.length() >= n
    puts string[n - 1]
    break
  end
end
