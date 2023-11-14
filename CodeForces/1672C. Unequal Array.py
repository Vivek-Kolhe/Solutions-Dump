for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  maxi, mini = -1, -1
  for i in range(n - 1):
    if arr[i] == arr[i+1]:
      if mini == -1:
        mini = i
      maxi = i
  if maxi == mini:
    print(0)
  else:
    print(max(1, maxi - mini - 1))
