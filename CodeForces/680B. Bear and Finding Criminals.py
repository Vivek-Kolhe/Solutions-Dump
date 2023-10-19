n, pos = map(int, input().split())
arr = list(map(int, input().split()))
if n == 1:
  if arr[0] == 1:
    print(1)
  else:
    print(0)
else:
  pos -= 1
  cnt = 0
  l, r = pos, n - 1 - pos
  part = arr[:pos+l+1] if l < r else arr[pos-r:]
  for i in range((len(part) // 2) + 1):
    if part[i] == 1 and part[-1-i] == 1:
      if len(part) % 2 == 1 and i == len(part) // 2:
        cnt += 1
      else:
        cnt += 2
  cnt += sum(arr[pos+l+1:] if l < r else arr[:pos-r])
  print(cnt)
