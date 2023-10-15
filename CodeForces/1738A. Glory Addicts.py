for _ in range(int(input())):
  n = int(input())
  types = list(map(int, input().split()))
  dmg = list(map(int, input().split()))
  b0 = sorted([dmg[i] for i in range(n) if not types[i]], reverse=True)
  b1 = sorted([dmg[i] for i in range(n) if types[i]], reverse=True)
  if len(b0) == len(b1):
    print(2 * sum(dmg) - min(dmg))
  else:
    k = min(len(b0), len(b1))
    print(2 * (sum(b0[:k]) + sum(b1[:k])) + sum(b0[k:]) + sum(b1[k:]))
