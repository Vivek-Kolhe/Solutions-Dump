def pmat(n, m, mat):
  for i in range(n):
    for j in range(m):
      print(mat[i][j], end=" ")
    print()

for _ in range(int(input())):
  n, m = map(int, input().split())
  combs = [[1, 0], [0, 1]]
  if n == 2 and m == 2:
    pmat(n, m, combs)
  else:
    mat = []
    for i in range(2):
      temp = []
      for j in range(m // 2):
        if j % 2 == 0:
          temp += combs[i%2]
        else:
          temp += combs[i%2][::-1]
      mat.append(temp)
    for i in range(n - 2):
      if i % 2 == 0:
        mat.append(mat[-1])
      else:
        mat.append(mat[0])
    pmat(n, m, mat)
