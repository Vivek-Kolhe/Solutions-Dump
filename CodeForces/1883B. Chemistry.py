from collections import Counter

for _ in range(int(input())):
  n, k = map(int, input().split())
  s = input()
  c = Counter(s)
  odd = 0
  for key in c:
    if c[key] % 2 == 1:
      odd += 1
  print("NO" if k + 1 < odd else "YES")
