n, k = map(int, input().split())
arr = list(map(int, input().split()))
new = sorted(arr)
h = {}
pref = [new[0]]
for i in range(1, n):
  pref.insert(i, new[i] + pref[-1])
for i in range(n):
  if h.get(arr[i]) is None:
    h[arr[i]] = []
  h[arr[i]].append(i)
lim = -1
for i in range(n):
  if pref[i] <= k:
    lim = max(lim, i)
if lim == -1:
  print(0)
else:
  res = []
  for i in range(lim + 1):
    res.append(h[new[i]].pop(0) + 1)
  print(len(res))
  print(*res)
