for _ in range(int(input())):
  n, t = map(int, input().split())
  dur = list(map(int, input().split()))
  ent = list(map(int, input().split()))
  ans = -2
  for i in range(n):
    if i + dur[i] <= t and (ans == -2 or ent[ans] < ent[i]):
      ans = i
  print(ans + 1)
