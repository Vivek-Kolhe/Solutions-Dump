n = int(input())
arr = list(map(int, input().split()))
i = 0
while n > 0:
  n -= arr[i]
  i += 1
  if i > 6:
    if n < 1:
      break
    i = 0
print(i)
