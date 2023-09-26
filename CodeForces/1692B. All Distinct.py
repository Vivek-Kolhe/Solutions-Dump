for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  dis = list(set(arr))
  if n % 2 == len(dis) % 2:
    print(len(dis))
  else:
    print(len(dis) - 1)
