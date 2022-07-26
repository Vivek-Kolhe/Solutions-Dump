n = int(input())
count, i = 0, 5
while True:
    if i > n:
        break
    count += n // i
    i *= 5
print(count)