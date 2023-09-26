n = int(input())
table = []
table.append([1] * n)
for i in range(n - 1):
  sum_ = 0
  temp = []
  for j in table[i]:
    sum_ += j
    temp.append(sum_)
  table.append(temp)
print(table[-1][-1])
