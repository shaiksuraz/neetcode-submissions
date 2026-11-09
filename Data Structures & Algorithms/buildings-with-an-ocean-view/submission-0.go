func findBuildings(heights []int) []int {
	n := len(heights)
	maxHeightTillNow := heights[n - 1]
	stack := []int{n-1}
	for idx := n-2; idx>=0; idx-- {
		if heights[idx] > maxHeightTillNow {
			stack = append(stack, idx)
			maxHeightTillNow = heights[idx]
		}
	}
	i, j := 0 ,len(stack)-1
	for i < j {
		stack[i], stack[j] = stack[j], stack[i]
		i++
		j--
	}
	return stack
}
