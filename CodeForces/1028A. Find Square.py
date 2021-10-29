n, m = map(int, input().split())
arr = []
for _ in range(n):
  arr.append(input())
for i in range(n):
  if "B" in arr[i]:
    count, index, ini_row = arr[i].count("B"), arr[i].find("B"), i
    break
col = index + ((count // 2) + 1)
row = ini_row + ((count // 2) + 1)
print(row, col)
