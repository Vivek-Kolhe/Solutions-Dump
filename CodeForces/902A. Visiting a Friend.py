n, m = map(int, input().split())
tps = {}
for _ in range(n):
  a, b = map(int, input().split())
  if tps.get(a) is None:
    tps[a] = b
  else:
    tps[a] = max(b, tps[a])
keys = list(tps.keys())
if keys[0] != 0:
  print("NO")
  exit(0)
canbe = 0
for i in range(len(keys)):
  if keys[i] <= canbe <= tps[keys[i]]:
    canbe = max(canbe, tps[keys[i]])
    if canbe >= m:
      print("YES")
      break
else:
  print("NO")
