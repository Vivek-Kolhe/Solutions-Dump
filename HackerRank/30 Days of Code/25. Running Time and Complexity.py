# i used this approach, doesn't work for 2 test cases (time limit exceeded).
n = int(input())
for _ in range(n):
    num = int(input())
    flag = 0
    if num == 1:
        print("Not prime")
    else:
        for i in range(2, num):
            if num % i == 0:
                print("Not prime")
                flag = 1
                break
        if not flag:
            print("Prime")

# Referred to GFG for this.
def isprime(num):
    if num <= 1:
        return "Not prime"
    if num <= 3:
        return "Prime"
    
    if (num % 2 == 0) or (num % 3 == 0):
        return "Not prime"
    
    i = 5
    while(i * i <= num):
        if (num % i == 0) or (num % (i + 2) == 0):
            return "Not prime"
        i += 6
    return "Prime"

n = int(input())
for _ in range(n):
    num = int(input())
    print(isprime(num))
