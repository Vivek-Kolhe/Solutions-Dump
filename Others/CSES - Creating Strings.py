from itertools import permutations
n = sorted(input())
arr = set(permutations(n))
result = []
print(len(arr))
for item in arr:
    result.append("".join(item))
result.sort()
for item in result:
    print(item)