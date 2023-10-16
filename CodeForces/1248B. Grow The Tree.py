n = int(input())
arr = list(map(int, input().split()))
arr.sort()
mid = n // 2
one, two = arr[:mid], arr[mid:]
print(sum(one) ** 2 + sum(two) ** 2)
