for _ in range(int(input())):
  n = int(input())
  candies = list(map(int, input().split()))
  one, two = candies.count(1), candies.count(2)
  if sum(candies) % 2 != 0:
    print("NO")
  else:
    sum_ = sum(candies) // 2
    if sum_ % 2 == 0 or sum_ % 2 != 0 and one != 0:
      print("YES")
    else:
      print("NO")
