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
        int l, r;
        cin >> l >> r;
        l--;
        r--;
        cout << query(1, 0, n - 1, l, r) << endl;
        m--;
    }
}
 
int main() {
    solve();
    return 0;
}
