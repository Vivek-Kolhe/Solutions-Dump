n = input()
a, b = int(n[:-1]), int(n[:-2] + n[-1])
arr = [int(n), a, b]
print(max(arr))
