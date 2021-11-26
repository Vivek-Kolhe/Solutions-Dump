for _ in range(int(input())):
  n = input()
  flag = False
  if int(n) % 2 == 0:
    print(0)
  elif int(n[0]) % 2 == 0:
    print(1)
  else:
    odd_count = 0
    for i in range(len(n)):
      if int(n[i]) % 2 != 0:
        odd_count += 1
    if odd_count == len(n):
      print(-1)
    else:
      print(2)
