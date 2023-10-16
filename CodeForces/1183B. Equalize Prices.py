for _ in range(int(input())):
  n, k = map(int, input().split())
  arr = list(map(int, input().split()))
  maxi, mini = max(arr), min(arr)
  if maxi - mini > 2 * k:
    print(-1)
  else:
    print(mini + k)
