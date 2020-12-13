package mathUnit

// for implementation accumulation
func GetSum(n int) (sum int) {
	for i := 0; i <= n; i++ {
		sum += i
	}
	sum += 1		// 假装出错
	return
}

// recursive implementation accumulation
func GetSumRecursive(n int) int {
	// termination condition 终止条件
	if n == 1 {
		return 1
	}
	return n + GetSumRecursive(n-1)
}
