n, m, a = map(int, input().split())
col = n // a if n % a == 0 else (n // a) + 1
row = m // a if m % a == 0 else (m // a) + 1
print(col * row)
