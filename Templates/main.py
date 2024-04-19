import sys

class SegTree:
    def __init__(self, n):
        self.seg = [sys.maxsize] *  (4 * n)

    def build(self, idx, low, high, arr):
        if low == high:
            self.seg[idx] = arr[low]
            return
        mid = (low + high) >> 1
        self.build(2 * idx, low, mid, arr)
        self.build(2 * idx + 1, mid + 1, high, arr)
        self.seg[idx] = min(self.seg[2*idx], self.seg[2*idx+1])
    
    def update(self, idx, low, high, pos, val):
        if pos < low or pos > high:
            return
        if low == high:
            self.seg[idx] = val
            return
        mid = (low + high) >> 1
        self.update(2 * idx, low, mid, pos, val)
        self.update(2 * idx + 1, mid + 1, high, pos, val)
        self.seg[idx] = min(self.seg[2*idx], self.seg[2*idx+1])

    def query(self, idx, low, high, l, r):
        if l > high or r < low:
            return sys.maxsize
        if l <= low and high <= r:
            return self.seg[idx]
        mid = (low + high) >> 1
        return min(self.query(2 * idx, low, mid, l, r), self.query(2 * idx + 1, mid + 1, high, l, r))
