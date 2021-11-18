for _ in range(int(input())):
  arr = list(map(int, input().split()))
  final_1 = arr[1] if arr[1] > arr[0] else arr[0]
  final_2 = arr[2] if arr[2] > arr[3] else arr[3]
  arr.sort()
  if (arr.index(final_1) == 2 or arr.index(final_1) == 3) and (arr.index(final_2) == 2 or arr.index(final_2) == 3):
    print("YES")
  else:
    print("NO")
