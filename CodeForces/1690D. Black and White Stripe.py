for _ in range(int(input())):
  n, k = map(int, input().split())
  s = input()
  cnt = s[:k].count("W")
  ans = cnt
  for i in range(k, n):
    if s[i] == "W":
      cnt += 1
    if s[i-k] == "W":
      cnt -= 1
    ans = min(ans, cnt)
  print(ans)
