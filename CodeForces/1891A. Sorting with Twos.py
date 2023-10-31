def isSorted(arr):
  n = len(arr)
  if n == 0 or n == 1:
    return True
  for i in range(1, n):
    if arr[i-1] > arr[i]:
      return False
  return True

for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  if n <= 3:
    print("YES")
    continue
  powers = [1, 3, 7, 15]
  flag = True
  for i in range(3):
    if not powers[i] < n:
      break
    l = powers[i]
    r = powers[i+1] if powers[i+1] < n else n - 1
    if not isSorted(arr[l+1:r+1]):
      print("NO")
      flag = False
      break
  if flag:
    if not isSorted(arr[r+1:]):
      print("NO")
      continue
    print("YES")
