func numberOfSteps(num int) int {
	var step int = 0
	for num != 0 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num - 1
		}
		step = step + 1
	}
	return step
}