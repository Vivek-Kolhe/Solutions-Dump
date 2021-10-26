n, b, d = map(int, input().split())
arr = list(map(int, input().split()))
sum_, count = 0, 0
for i in range(n):
  if arr[i] <= b:
    sum_ += arr[i]
    if sum_ > d:
      count += 1
      sum_ = 0
print(count)
