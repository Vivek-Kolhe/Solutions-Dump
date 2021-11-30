vowels = ["a", "e", "i", "o", "u", "y"]
string = gets.chomp.downcase.gsub(" ", "")
if vowels.include?string[-2]
  puts "YES"
else
  puts "NO"
end
