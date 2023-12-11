n, m = map(int, input().split())
s = list(input())
mp = {}
for i in range(n):
  if mp.get(s[i]) is None:
    mp[s[i]] = []
  mp[s[i]].append(i)
for _ in range(m):
  x, y = input().split()
  if mp.get(x) is None:
    mp[x] = []
  if mp.get(y) is None:
    mp[y] = []
  mp[x], mp[y] = mp[y], mp[x]
res = [""] * n
for key in mp:
  for ele in mp[key]:
    res[ele] = key
print("".join(res))
