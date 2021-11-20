for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  max_ = 0
  for i in range(n - 1):
    if max_ < arr[i] * arr[i + 1]:
      max_ = arr[i] * arr[i + 1]
  print(max_)
