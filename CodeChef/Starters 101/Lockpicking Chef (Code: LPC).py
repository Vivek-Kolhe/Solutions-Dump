# cook your dish here
def close(n, target):
    if n > target:
        return min(abs(n - target), abs(10 - n + target))
    return min(abs(n - target), abs(10 + n - target))

for _ in range(int(input())):
    n, m = map(int, input().split())
    s = input()
    key = input()
    res = []
    for i in range(n - m + 1):
        temp = s[i:i+m]
        count = 0
        for j in range(m):
            count += close(int(temp[j]), int(key[j]))
        res.append(count)
    print(min(res))
