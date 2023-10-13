for _ in range(int(input())):
  n = int(input())
  bigrams = input().split()
  res = bigrams[0]
  i = 0
  while i < len(bigrams) - 1:
    if bigrams[i][-1] != bigrams[i+1][0]:
      bigrams.insert(i + 1, bigrams[i][-1] + bigrams[i+1][0])
      break
    i += 1
  for i in range(1, len(bigrams)):
    res += bigrams[i][-1]
  print(res if len(res) == n else res + "a")
