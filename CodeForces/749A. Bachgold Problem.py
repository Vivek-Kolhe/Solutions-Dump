n = int(input())
result = [2 for i in range(n // 2)]
if n % 2 != 0:
  result[-1] = 3
print(len(result))
print(*result)
