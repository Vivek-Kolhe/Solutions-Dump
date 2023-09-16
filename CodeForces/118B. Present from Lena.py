n = int(input())
max_len = (n * 2) + 1
ini = []
for i in range(n + 1):
  temp = []
  for j in range(i + 1):
    temp.append(j)
  if len(temp) > 1:
    temp.extend(temp[-2::-1])
  ini.append(temp)
ini.extend(ini[-2::-1])
for i in range(len(ini)):
  temp = (max_len - len(ini[i])) // 2
  ini[i] = [" "] * temp + ini[i]
for item in ini:
  print(*item)
