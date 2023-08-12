lo, lc, ro, rc = 0, 0, 0, 0
for _ in range(int(input())):
  l, r = map(int, input().split())
  if l == 0:
    lc += 1
  else:
    lo += 1
  if r == 0:
    rc += 1
  else:
    ro += 1
print(min(lc, lo) + min(rc, ro))
