package problems

func distanceTraveled(mainTank int, additionalTank int) int {
	if mainTank < 5 || additionalTank == 0 {
		return mainTank * 10
	}
	b5 := mainTank / 5
	move := min(b5, additionalTank)
	return 50*b5 + distanceTraveled(mainTank-(b5*5)+move, additionalTank-move)
}
