for _ in range(int(input())):
  n = int(input())
  arr = sorted(list(map(int, input().split())))
  if len(arr) == 1:
    print(arr[0])
  else:
    max_ = arr[0]
    for i in range(n - 1):
      max_ = max(max_, arr[i + 1] - arr[i])
    print(max_)
