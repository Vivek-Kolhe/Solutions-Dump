# Enter your code here. Read input from STDIN. Print output to STDOUT

def check_float(n):
    try:
        if "." not in n:
            return False
        n = float(n)
        return True
    except:
        return False
    
t = int(input())
for _ in range(t):
    print(check_float(input()))
