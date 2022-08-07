package main

func main() {

}
func climbStairs(n int) int {
	res := make([]int, n)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			res[i] = 1
			continue
		}

		if i == 1 {
			res[i] = 2
			continue
		}

		res[i] = res[i-1] + res[i-2]
	}

	return res[n-1]
}
