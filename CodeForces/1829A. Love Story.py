for _ in range(int(input())):
  s = 'codeforces'
  string = input()
  count = 0
  for i in range(10):
    if string[i] != s[i]:
      count += 1
  print(count)
