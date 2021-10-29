n = int(input())
chessboard = []
count = 0
for i in range(n):
  temp = ["."] * n
  if i % 2 == 0:
    for j in range(0, n, 2):
      count += 1
      temp[j] = "C"
  else:
    for k in range(1, n, 2):
      count += 1
      temp[k] = "C"
  chessboard.append(temp)
print(count)
for i in range(n):
  print("".join(chessboard[i]))
