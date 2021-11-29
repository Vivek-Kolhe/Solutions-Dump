for _ in range(int(input())):
  n = int(input())
  arr = []
  arr.append("()" * n)
  for i in range(n):
    temp = "(" * (n - i) + ")" * (n - i) + "()" * i
    arr.append(temp)
  for i in range(n):
    print(arr[i])
