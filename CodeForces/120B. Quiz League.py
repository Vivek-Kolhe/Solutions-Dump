import sys

sys.stdin = open("input.txt", "r")
sys.stdout= open("output.txt", "w")
n, k = map(int, input().split())
arr = list(map(int, input().split()))
i = k - 1
while not arr[i]:
  i += 1
  if i > n - 1:
    i = 0
print(i + 1)
