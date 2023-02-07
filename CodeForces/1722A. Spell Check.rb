for _ in 0...gets.chomp.to_i
  n = gets.chomp.to_i
  string = gets.chomp.split('')
  if (string.length() > 5) or (string.length() < 5)
    puts 'NO'
  else
    perms = 'Timur'.split('').permutation().to_a
    if perms.include?string
      puts 'YES'
    else
      puts 'NO'
    end
  end
end
