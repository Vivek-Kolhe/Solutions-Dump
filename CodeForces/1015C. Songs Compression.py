n, m = map(int, input().split())
tot = 0
arr = []
for _ in range(n):
  a, b = map(int, input().split())
  tot += a
  arr.append([a, b])
diff = [0] * n
for i in range(n):
  diff[i] = arr[i][0] - arr[i][1]
diff.sort(reverse=True)
for i in range(n):
  if tot <= m:
    print(i)
    break
  tot -= diff[i]
else:
  if tot <= m:
    print(n)
  else:
    print(-1)
