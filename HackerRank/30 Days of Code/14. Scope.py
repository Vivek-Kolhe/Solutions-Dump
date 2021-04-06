class Difference:
    def __init__(self, a):
        self.__elements = a
        self.maximumDifference = 0

    def computeDifference(self):
        l = []
        for i in self.__elements:
            for j in self.__elements:
                l.append(abs(i - j))
        self.maximumDifference = max(set(l))
        return self.maximumDifference

_ = input()
a = [int(e) for e in input().split(' ')]

d = Difference(a)
d.computeDifference()

print(d.maximumDifference)
