for _ in range(int(input())):
  ticket = input()
  count_1, count_2 = 0, 0
  for i in range(3):
    count_1 += int(ticket[i])
    count_2 += int(ticket[5 - i])
  if count_1 == count_2:
    print("YES")
  else:
    print("NO")
