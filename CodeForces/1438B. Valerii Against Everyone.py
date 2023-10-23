for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  new = list(set(arr))
  if len(arr) == len(new):
    print("NO")
  else:
    print("YES")
