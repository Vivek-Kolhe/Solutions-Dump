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
    int pos;
    cin >> pos;
    pos = pos - 1;
    arr.erase(arr.begin() + pos);
    int start, end;
    cin >> start >> end;
    start = start - 1;
    end = end - 1;
    arr.erase(arr.begin() + start, arr.begin() + end);
    int size = arr.size();
    cout << size << endl;
    for(int i = 0; i < size; i++)
    {
        cout << arr[i] << " ";
    }
    return 0;
}
