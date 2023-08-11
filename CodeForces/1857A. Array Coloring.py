for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  if len(arr) == 1:
    print("NO")
  else:
    odd, even = 0, 0
    for i in range(n):
      if arr[i] % 2 == 0:
        even += 1
      else:
        odd += 1
    if odd % 2 == 0:
      print("YES")
    else:
      print("NO")
