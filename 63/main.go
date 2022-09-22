package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	// 状态是机器人所在位置，选择只有向下或向右
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// base case：初始化第一行和第一列, 遇到障碍物, 后面的全部置为0
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue // 默认 dp[i][j]=0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	println(uniquePathsWithObstacles(obstacleGrid))
}
