def score(p, t):
  return max((3 * p) / 10, p - ((p / 250) * t))

a, b, c, d = map(int, input().split())
misha, vasya = score(a, c), score(b, d)
if misha > vasya:
  print('Misha')
elif misha < vasya:
  print('Vasya')
else:
  print('Tie')
