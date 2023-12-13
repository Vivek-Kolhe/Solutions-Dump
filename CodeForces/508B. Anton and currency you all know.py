s = list(input())
ans = -1
n = len(s)
for i in range(n - 1):
  if int(s[i]) & 1 == 0:
    ans = i
    if int(s[n-1]) > int(s[i]):
      break
if ans == -1:
  print(ans)
else:
  s[ans], s[n-1] = s[n-1], s[ans]
  print("".join(s))
