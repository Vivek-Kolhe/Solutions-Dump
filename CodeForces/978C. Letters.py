n, m = map(int, input().split())
dorms = list(map(int, input().split()))
letters = list(map(int, input().split()))
dorms.insert(0, 0)
s = 0
for i in range(n + 1):
  store = dorms[i]
  dorms[i] += s
  s += store
print(dorms)
for letter in letters:
  ans = -1
  low, high = 0, n
  while low <= high:
    mid = (low + high) // 2
    if letter <= dorms[mid]:
      ans = mid
      high = mid - 1
    else:
      low = mid + 1
  print(ans, letter - dorms[ans-1])
