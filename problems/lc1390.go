package problems

func sumDiv(x int) int {
	if x == 1 {
		return 0
	}
	sum := 1 + x
	cnt := 2
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			cnt++
			sum += i
			if i*i < x {
				cnt++
				sum += x / i
			}
			if cnt > 4 {
				break
			}
		}
	}
	if cnt == 4 {
		return sum
	}
	return 0
}
func sumFourDivisors(nums []int) int {
	ans := 0
	for _, x := range nums {
		ans += sumDiv(x)
	}
	return ans
}
