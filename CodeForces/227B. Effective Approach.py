n = int(input())
arr = list(map(int, input().split()))
m = int(input())
queries = list(map(int, input().split()))
v, p = 0, 0
new = [0] * (10 ** 5 + 3)
for i in range(n):
  new[arr[i]] = i
for i in range(m):
  temp = new[queries[i]]
  v += temp + 1
  p += n - temp
print(v, p)
