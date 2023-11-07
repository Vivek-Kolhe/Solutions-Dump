n, k = map(int, input().split())
cnt = 0
for _ in range(n):
  num = list(map(int, list(input())))
  num.sort()
  for i in range(k + 1):
    if i not in num:
      break
  else:
    cnt += 1
print(cnt)
