x, y, z, t1, t2, t3 = map(int, input().split())
stairs = abs(x - y) * t1
ele = (3 * t3) + (abs(x - z) * t2) + (abs(x - y) * t2)
if ele <= stairs:
  print("YES")
else:
  print("NO")
