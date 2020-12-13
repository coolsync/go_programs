package mathUnit

//获取斐波那契数列第n项的非递归实现
//1,1,2,3,5,8,13,21,34,55
func GetFobnacci(n int) int {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

//获取斐波那契数列第n项的递归实现
//1,1,2,3,5,8,13,21,34,55
func GetFobnacciRecursive(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return GetFobnacciRecursive(n - 1) + GetFobnacciRecursive(n - 2)
}