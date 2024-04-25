#include <iostream>
#include <vector>
using namespace std;

void dfs(int n, int m, int row, int col, vector<vector<char>>& grid, vector<vector<bool>>& visited) {
    visited[row][col] = true;
    int delRows[] = {0, 0, 1, -1};
    int delCols[] = {1, -1, 0, 0};
    for (int i = 0; i < 4; ++i) {
        int nrow = row + delRows[i];
        int ncol = col + delCols[i];
        if (nrow >= 0 && nrow < n && ncol >= 0 && ncol < m && grid[nrow][ncol] == '.' && !visited[nrow][ncol]) {
            dfs(n, m, nrow, ncol, grid, visited);
        }
    }
}

int main() {
    int n, m;
    cin >> n >> m;
    vector<vector<char>> grid(n, vector<char>(m));
    for (int i = 0; i < n; ++i) {
        for (int j = 0; j < m; ++j) {
            cin >> grid[i][j];
        }
    }
    vector<vector<bool>> visited(n, vector<bool>(m, false));
    int ans = 0;
    for (int i = 0; i < n; ++i) {
        for (int j = 0; j < m; ++j) {
            if (grid[i][j] == '.' && !visited[i][j]) {
                ++ans;
                dfs(n, m, i, j, grid, visited);
            }
        }
    }
    cout << ans << endl;
    return 0;
}
