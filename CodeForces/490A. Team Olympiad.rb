n = gets.chomp.to_i
arr = gets.chomp.split().map{|x| x.to_i}
count_a, count_b, count_c = 0, 0, 0
ind_a, ind_b, ind_c = [], [], []
for i in 0...n
  if arr[i] == 1
    count_a += 1
    ind_a.append(i)
  elsif arr[i] == 2
    count_b += 1
    ind_b.append(i)
  else
    count_c += 1
    ind_c.append(i)
  end
end
poss_teams = [count_a, count_b, count_c].min(1)
puts poss_teams[0]
for i in 0...poss_teams[0]
  puts "#{ind_a[i] + 1} #{ind_b[i] + 1} #{ind_c[i] + 1}"
end
