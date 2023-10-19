c, v0, v1, a, l = map(int, input().split())
cnt = 0
while c > 0:
  if cnt == 0:
    if v0 < v1:
      c -= v0
    else:
      c -= v1
    cnt += 1
    continue
  c += l
  toread = v0 + (cnt * a)
  if toread < v1:
    c -= toread
  else:
    c-= v1
  cnt += 1
print(cnt)
