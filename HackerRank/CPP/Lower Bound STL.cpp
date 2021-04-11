#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;

int main()
{
    int n;
    cin >> n;
    vector<int> arr;
    for(int i = 0; i < n; i++)
    {
        int temp;
        cin >> temp;
        arr.push_back((temp));
    }
    int queries;
    cin >> queries;
    for(int i = 0; i < queries; i++)
    {
        int query;
        cin >> query;
        vector<int>::iterator it = lower_bound(arr.begin(), arr.end(), query);
        if(arr[it - arr.begin()] == query)
        {
            cout << "Yes " << it - arr.begin() + 1 <<endl;
        }
        else
        {
            cout << "No " << it - arr.begin() + 1 <<endl;
        }
    }
    return 0;
}
