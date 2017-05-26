package fibo

var Result int64

func Fibonacci(n int) int64 {
	iterations := make([]int64, n+1)
	if n == 0 {
		Result = 0
		return 0
	} else if n == 1 {
		Result = 0
		return 1
	}
	iterations[0] = 0
	iterations[1] = 1

	for i := 2; i <= n; i++ {
		iterations[i] = iterations[i-1] + iterations[i-2]

	}
	Result = iterations[n-1]
	return iterations[n]
}
