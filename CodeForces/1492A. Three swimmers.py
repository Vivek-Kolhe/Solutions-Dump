def ceil(p, a):
  if p % a == 0:
    return p // a
  return (p // a) + 1

for _ in range(int(input())):
  p, a, b, c = map(int, input().split())
  ceils = [ceil(p, a), ceil(p, b), ceil(p, c)]
  print(min([a * ceils[0] - p, b * ceils[1] - p, c * ceils[2] - p]))
