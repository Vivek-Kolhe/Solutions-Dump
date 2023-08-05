for _ in range(int(input())):
  pi = '314159265358979323846264338327'
  num = input()
  count = 0
  for i in range(len(num)):
    if pi[i] == num[i]:
      count += 1
    else:
      break
  print(count)
