require "prime"

def palindrome?(string)
  if string == string.reverse
    return true
  end
  return false
end

result = -> (n) do
  Prime.each.lazy.select {
    |x| palindrome?(x.to_s)
  }.first(n)
end

n = gets.chomp.to_i
p result.(n)
