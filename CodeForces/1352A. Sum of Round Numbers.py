for _ in range(int(input())):
  n = input()
  result = []
  for i in range(len(n)):
    temp = n[i] + ("0" * (len(n) - i - 1))
    if int(temp) != 0:
      result.append(temp)
  print(len(result))
  for i in range(len(result)):
    print(result[i], end = " ")
  print()
