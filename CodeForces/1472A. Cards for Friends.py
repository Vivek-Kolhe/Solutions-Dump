for _ in range(int(input())):
  w, h, n = map(int, input().split())
  count = 1
  while w % 2 == 0:
    w = w // 2
    count = count * 2
  while h % 2 == 0:
    h = h // 2
    count = count * 2
  if count >= n:
    print("yes")
  else:
    print("no")
