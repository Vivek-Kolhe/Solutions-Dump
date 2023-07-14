for _ in range(int(input())):
  n = int(input())
  sum = (n * (n + 1)) // 2
  pow2, ind = 0, 0
  while True:
    temp = 2 ** ind
    if temp > n:
      break
    pow2 += 2 ** ind
    ind += 1
  print(sum - (2 * pow2))
