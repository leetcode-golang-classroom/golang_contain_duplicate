package sol

func containsDuplicate(nums []int) bool {
	hash := make(map[int]struct{})
	for _, val := range nums {
		if _, exist := hash[val]; exist {
			return true
		}
		hash[val] = struct{}{}
	}
	return false
}
