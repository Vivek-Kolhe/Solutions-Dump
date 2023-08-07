n = int(input())
arr = list(map(int, input().split()))
ind = []
for i in range(1, n + 1):
  ind.append(arr.index(i) + 1)
print(*ind)
