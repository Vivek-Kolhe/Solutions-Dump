n = int(input())
arr = []
for _ in range(n):
  arr.append(int(input()))
bi = [i // 2 for i in arr]
s = sum(bi)
for i in range(n):
  if arr[i] % 2 != 0:
    bi[i] += 1
    s += 1
  if s == 0:
    break
for ele in bi:
  print(ele)
