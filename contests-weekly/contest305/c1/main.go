package main

func main() {

}

func arithmeticTriplets(nums []int, diff int) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				diffJI := nums[j] - nums[i]
				diffKJ := nums[k] - nums[j]

				if diffJI == diff && diffKJ == diff {
					count++
				}
			}
		}
	}
	return count
}
