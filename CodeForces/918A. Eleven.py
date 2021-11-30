n = int(input())
n1, n2 = 0, 1
fib = [0]
for i in range(n + 1):
    n3 = n1 + n2
    fib.append(n3)
    n2 = n1
    n1 = n3
name = ""
for i in range(1, n + 1):
    if i in fib:
        name += "O"
    else:
        name += "o"
print(name)
