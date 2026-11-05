func maxArea(heights []int) int {
	var res int
	i := 0
	j := len(heights) - 1
	for i < j {
		var currentArea int
		k := min(heights[i], heights[j])
		currentArea = k * (j-i)
		res = max(res, currentArea)
		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}
	return res
}
