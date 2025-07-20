package piscine

func CanJump(nums []uint) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}
	currentIndex := 0
	lastIndex := len(nums) - 1
	for {
		steps := nums[currentIndex]
		if steps == 0 && currentIndex != lastIndex {
			return false
		}
		nextIndex := currentIndex + int(steps)
		if nextIndex == lastIndex {
			return true
		}
		if nextIndex > lastIndex {
			return false
		}
		currentIndex = nextIndex
	}
}