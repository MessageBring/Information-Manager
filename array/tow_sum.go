package array

func TwoSum(nums []int, target int) (res []int) {
	var tempMap = map[int]int{}
	for i, v := range nums {
		if _, ok := tempMap[nums[i]]; ok {
			res = append(res, tempMap[v], i)
			return
		}
		temp := target - v
		tempMap[temp] = i
	}
	return
}
