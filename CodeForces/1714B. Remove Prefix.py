for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))[::-1]
  h = {}
  if n == len(set(arr)):
    print(0)
  else:
    for i in range(n):
      if h.get(arr[i]):
        print(n - i)
        break
      else:
        h[arr[i]] = True
