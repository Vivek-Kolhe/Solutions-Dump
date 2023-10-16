for _ in range(int(input())):
  arr = list(map(int, input().split()))
  arr.sort()
  maxi = arr[-1]
  if maxi - (arr[0] + arr[1]) <= 1:
    print("Yes")
  else:
    print("No")
