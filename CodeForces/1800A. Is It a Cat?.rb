gets.chomp.to_i.times do
  n = gets.chomp.to_i
  string = gets.chomp
  if (string.start_with?('M') or string.start_with?('m')) and (string.end_with?('W') or string.end_with?('w'))
    meow = string[0].downcase
    for i in 1...n
      if meow[-1].downcase != string[i].downcase
        meow += string[i].downcase
      end
    end
    if meow == "meow"
      puts 'YES'
    else
      puts 'NO'
    end
  else
    puts 'NO'
  end
end
