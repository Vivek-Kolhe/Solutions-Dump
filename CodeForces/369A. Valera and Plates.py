n, b, p = map(int, input().split())
arr = list(map(int, input().split()))
count = 0
for i in range(n):
  if arr[i] == 1:
    if b > 0:
      count += 1
      b -= 1
  else:
    if p > 0:
      p -= 1
      count += 1
    else:
      if b > 0:
        b -= 1
        count += 1
print(n - count)
