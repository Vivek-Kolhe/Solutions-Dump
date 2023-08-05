for _ in range(int(input())):
  n, d = map(int, input().split())
  num = input()
  i = 0
  while i < n and int(num[i]) >= d:
    i += 1
  print(num[:i] + str(d) + num[i:])
