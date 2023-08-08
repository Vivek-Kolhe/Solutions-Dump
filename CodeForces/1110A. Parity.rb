b, k = gets.chomp.split.map(&:to_i)
arr = gets.chomp.split.map(&:to_i)
if b.even?
  if arr[-1].even?
    puts "even"
  else
    puts "odd"
  end
else
  count = 0
  for i in arr
    if i.odd?
      count += 1
    end
  end
  if count.even?
    puts "even"
  else
    puts "odd"
  end
end
