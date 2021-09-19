n = int(input())
arr = list(map(int, input().split()))
max_ =  0
for i in range(3):
    temp = arr.count(i + 1)
    if temp > max_:
        max_ = temp
print(n - max_)
