n, s = map(int, input().split())
h = {}
for _ in range(n):
  f, p = map(int, input().split())
  if h.get(f) is None:
    h[f] = []
  h[f].append(p)
arr = sorted(h.keys())[::-1]
wait, ind = 0, 0
for i in range(s, -1, -1):
  wait += 0 if s == i else 1
  if h.get(i) is not None:
    maxi = max(h[i])
    if maxi != 0:
      maxi -= wait
    wait += max(0, maxi)
    ind += 1
print(wait)
