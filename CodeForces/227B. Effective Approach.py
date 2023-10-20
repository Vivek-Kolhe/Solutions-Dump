n = int(input())
arr = list(map(int, input().split()))
m = int(input())
queries = list(map(int, input().split()))
v, p = 0, 0
h = {}
for i in range(n):
  h[arr[i]] = i
for i in range(m):
  temp = h[queries[i]]
  v += temp + 1
  p += n - temp
print(v, p)
