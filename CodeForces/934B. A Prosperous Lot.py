n = int(input())
if n < 1:
  print("1")
else:
  if n >= pow(10, 6):
    print(-1)
    exit(0)
  ans = "8" * (n // 2)
  if n & 1:
    ans += "4"
  ans = int(ans)
  print(-1 if ans > pow(10, 18) else ans)
