for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  s = [arr[0]]
  cnt = 0
  for i in range(1, n):
    sl = len(s)
    while arr[i] < s[-1]:
      s.pop()
      cnt += 1
      sl -= 1
      if sl == 0:
        break
    s.append(arr[i])
  print(cnt)
