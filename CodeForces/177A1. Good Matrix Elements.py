n = int(input())
total = 0
matrix = [list(map(int, input().split())) for _ in range(n)]
mid = (n-1)//2
mid_el = matrix[mid][mid]
total += sum(matrix[mid])
for i in range(n):
    total += matrix[i][mid] + matrix[i][i] + matrix[n-1-i][i]
print(total - (3 * mid_el))
