n, m = map(int, input().split())
nip, cip = [], []
for _ in range(n):
  nip.append(input())
for _ in range(m):
  cip.append(input())
for name in nip:
  temp = name.split()
  for i in range(m):
    if cip[i].split()[-1][:-1] == temp[-1]:
      cip[i] += f" #{temp[0]}"
for c in cip:
  print(c)
