max_ = []
n, k = gets.chomp.split().map{|x| x.to_i}
for _ in 0...n
  f, t = gets.chomp.split().map{|x| x.to_i}
  if t <= k
    max_.append(f)
  else
    max_.append(f - (t - k))
  end
end
puts max_.max()
