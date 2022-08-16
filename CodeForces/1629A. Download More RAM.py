for _ in range(int(input())):
  n, k = map(int, input().split())
  a = list(map(int, input().split()))
  b = list(map(int, input().split()))
  new_arr = sorted([[i, j] for i, j in zip(a, b)], key = lambda x: x[0])
  for i in range(n):
    if new_arr[i][0] <= k:
      k += new_arr[i][-1]
  print(k)
