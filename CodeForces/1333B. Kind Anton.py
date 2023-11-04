for _ in range(int(input())):
  n = int(input())
  a = list(map(int, input().split()))
  b = list(map(int, input().split()))
  if a[0] != b[0]:
    print("NO")
  else:
    h = {
        -1 : False,
        0 : False,
        1 : False
      }
    for i in range(1, n):
      if not h[a[i-1]]:
        h[a[i-1]] = True
      if a[i] == b[i]:
        continue
      diff = b[i] - a[i]
      if diff > 0:
        if not h[1]:
          print("NO")
          break
      if diff < 0:
        if not h[-1]:
          print("NO")
          break
    else:
      print("YES")
