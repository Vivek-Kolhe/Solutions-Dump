for _ in range(int(input())):
  n = int(input())
  pile_1, pile_2 = 0, 0
  for i in range(1, n + 1):
    if i < (n // 2) or i == n:
      pile_1 += 2 ** i
    else:
      pile_2 += 2 ** i
  print(abs(pile_1 - pile_2))
