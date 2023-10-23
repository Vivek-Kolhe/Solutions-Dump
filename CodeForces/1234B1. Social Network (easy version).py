n, k = map(int, input().split())
arr = list(map(int, input().split()))
q = []
for i in range(n):
  if arr[i] not in q:
    q.insert(0, arr[i])
  if len(q) > k:
    q.pop(-1)
print(len(q))
print(*q)
