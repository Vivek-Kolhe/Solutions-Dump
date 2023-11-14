sizes = ["S", "M", "L", "XL", "XXL"]
def optimalsize(store, size, sizes):
  ind = sizes.index(size)
  l, r = ind - 1, ind + 1
  if store[size] > 0:
    store[size] -= 1
    return size
  while l > -1 or r < 5:
    if r < 5:
      if store[sizes[r]] > 0:
        store[sizes[r]] -= 1
        return sizes[r]
      r += 1
    if l > -1:
      if store[sizes[l]] > 0:
        store[sizes[l]] -= 1
        return sizes[l]
      l -= 1
      
s = list(map(int, input().split()))
store = {sizes[i] : s[i] for i in range(5)}
for _ in range(int(input())):
  size = input()
  print(optimalsize(store, size, sizes))
