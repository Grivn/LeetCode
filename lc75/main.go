package main

func main() {

}

func sortColors(nums []int)  {
	res := sort(nums)
	copy(nums, res)
}

func sort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	total := len(nums)

	middle := total/2

	left := nums[:middle]
	right := nums[middle:]

	return merge(sort(left), sort(right))
}

func merge(left, right []int) []int {
	var res []int

	for len(left) != 0 && len(right) != 0 {
		if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}

	if len(left) != 0 {
		res = append(res, left...)
	}

	if len(right) != 0 {
		res = append(res, right...)
	}

	return res
}
