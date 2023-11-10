import math
 
def is_prime(n):
  if n <= 1:
    return False
  for i in range(2, int(math.sqrt(n)) + 1):
    if n % i == 0:
      return False
  return True
  
for _ in range(int(input())):
  a, b = map(int, input().split())
  if a - b == 1 and is_prime(a + b):
    print("YES")
  else:
    print("NO")
