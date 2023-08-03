n = int(input())
cards = input()
if n // 11 > 0:
  pos = n // 11
  eights = cards.count('8')
  print(min(pos, eights))
else:
  print(0)
