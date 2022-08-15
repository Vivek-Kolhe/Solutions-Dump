for _ in range(int(input())):
  arr = list(map(int, input().split()))
  count = 0
  for i in range(1, 4):
    if arr[0] < arr[i]:
      count += 1
  print(count)
