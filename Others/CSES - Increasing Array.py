n = int(input())
arr = list(map(int, input().split()))
max_, result = 0, 0
for i in range(n):
    max_ = max(max_, arr[i])
    result += max_ - arr[i]
print(result)