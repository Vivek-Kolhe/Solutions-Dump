for _ in range(int(input())):
  x = input()
  sum = (int(x[0]) - 1) * 10
  for i in range(len(x)):
    sum += i + 1
  print(sum)
