n, a, b = map(int, input().split())
modulo = abs(b) % n
if b < 0:
  if modulo == 0: 
    print(a)
  else:
    ans = a + n - (abs(b) % n)
    print(ans if ans <= n else ans % n)
else:
  ans = a + (abs(b) % n)
  print(ans if ans <= n else ans % n)
