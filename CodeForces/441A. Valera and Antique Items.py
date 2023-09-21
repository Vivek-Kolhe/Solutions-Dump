n, purse = map(int, input().split())
res = []
for i in range(n):
  auc = list(map(int, input().split()))
  if min(auc[1:]) < purse:
    res.append(i + 1)
print(len(res))
if len(res) > 0:
  print(*res)
