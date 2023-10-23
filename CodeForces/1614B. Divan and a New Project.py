for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  new = sorted(arr)[::-1]
  ans = 0
  dist = 1
  h = {}
  for i in range(n):
    ans += 2 * abs(0 - dist) * new[i]
    if h.get(new[i]) is None:
      h[new[i]] = []
    l = h.get(new[i])
    if i % 2 == 0:
      l.append(dist)
    else:
      l.append(-dist)
      dist += 1
  res = [0]
  for ele in arr:
    val = h[ele]
    res.append(val.pop())
  print(ans)
  print(*res)
