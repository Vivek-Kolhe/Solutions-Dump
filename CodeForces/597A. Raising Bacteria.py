import math

def pow2(num):
  return int(pow(2, int(math.log(n, 2))))

n = int(input())
cnt = 0
while n > 0:
  n -= pow2(n)
  cnt += 1
print(cnt)
