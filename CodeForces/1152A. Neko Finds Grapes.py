def evenodd(arr):
  even, odd = 0, 0
  for ele in arr:
    if ele % 2 == 0: even += 1
    else: odd += 1
  return even, odd

n, m = map(int, input().split())
chests = list(map(int, input().split()))
keys = list(map(int, input().split()))
ec, oc = evenodd(chests)
ek, ok = evenodd(keys)
print(min(ec, ok) + min(oc, ek))
