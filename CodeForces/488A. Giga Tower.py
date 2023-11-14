n = int(input())
cnt = 1
for i in range(n + 1, 8888888889):
  s = str(i)
  if "8" in s:
    print(cnt)
    break
  else:
    cnt += 1
