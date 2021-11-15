a, b = map(int, input().split())
m, n = map(int, input().split())
arr_1 = list(map(int, input().split()))
arr_2 = list(map(int, input().split()))[::-1]
if arr_1[:m][-1] < arr_2[:n][-1]:
  print("YES")
else:
  print("NO")
