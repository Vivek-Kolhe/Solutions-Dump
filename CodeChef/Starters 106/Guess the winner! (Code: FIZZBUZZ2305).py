# cook your dish here
for _ in range(int(input())):
    n = int(input())
    if n <= 2 or n & 1 == 0:
        print("Bob")
    else:
        print("Alice")
