cube = lambda x: x ** 3 

def fibonacci(n):
    # return a list of fibonacci numbers
    seq = []
    for i in range(n):
        if i == 0:
            seq.append(0)
        elif i == 1:
            seq.append(1)
        else:
            seq.append(seq[i - 2] + seq[i - 1])
    return seq
    

if __name__ == '__main__':
    n = int(input())
    print(list(map(cube, fibonacci(n))))
