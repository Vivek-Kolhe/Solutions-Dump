dig = [8, 4, 2, 6]
n = int(input())
print(dig[(n%4)-1] if n > 0 else 1)
