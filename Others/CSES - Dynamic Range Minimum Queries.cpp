#include <bits/stdc++.h>
using namespace std;
 
#define ll long long int
 
ll arr[200005];
ll seg[4*200005];
 
void build(int idx, int low, int high) {
    if(low == high) {
        seg[idx] = arr[low];
        return;
    }
    int mid = (low + high) / 2;
    build(2 * idx, low, mid);
    build(2 * idx + 1, mid + 1, high);
    seg[idx] = min(seg[2*idx], seg[2*idx+1]);
}

void update(int idx, int low, int high, int pos, ll val) {
    if(pos < low || pos > high) {
        return;
    }
    if(low == high) {
        seg[idx] = val;
        return;
    }
    int mid = (low + high) / 2;
    update(2 * idx, low, mid, pos, val);
    update(2 * idx + 1, mid + 1, high, pos, val);
    seg[idx] = min(seg[2*idx], seg[2*idx+1]);
}

ll query(int idx, int low, int high, int l, int r) {
    if(low > r || high < l) {
        return (int)1e10;
    }
    if(l <= low && high <= r) {
        return seg[idx];
    }
    int mid = (low + high) / 2;
    return min(query(2 * idx, low, mid, l, r), query(2 * idx + 1, mid + 1, high, l, r));
}
 
void solve() {
    int n, m;
    cin >> n >> m;
    for(int i = 0; i < n; i++) {
        cin >> arr[i];
    }
    build(1, 0, n - 1);
    while(m > 0) {
        int ch, l;
        cin >> ch >> l;
        l--;
        if(ch == 1) {
            ll val;
            cin >> val;
            update(1, 0, n - 1, l, val);
        } else {
            int r;
            cin >> r;
            r--;
            cout << query(1, 0, n - 1, l, r) << endl;
        }
        m--;
    }
}
 
int main() {
    solve();
    return 0;
}
