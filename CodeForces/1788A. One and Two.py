for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  if arr.count(1) == n:
    print(1)
  elif arr.count(2) % 2 != 0:
    print(-1)
  else:
    half = arr.count(2) // 2
    count = 0
    for i in range(n):
      if arr[i] == 2:
        count += 1
      if count == half:
        print(i + 1)
        break
