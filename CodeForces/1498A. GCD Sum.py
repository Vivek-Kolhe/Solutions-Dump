import math
for _ in range(int(input())):
  n = int(input())
  ans = 0
  while True:
    digit_sum = 0
    for i in range(len(str(n))):
      digit_sum += int(str(n)[i])
    gcd_sum = math.gcd(n, digit_sum)
    if gcd_sum > 1:
      print(n)
      break
    else:
      n += 1
