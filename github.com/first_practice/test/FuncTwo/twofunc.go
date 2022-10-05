package functwo

func Sumintall(startat, endat int) int {
	var i int = startat
	var sum int = 0
	for i <= endat {
		sum += i
		i++
	}
	return sum
}

func Intdefer(startat, endat int) ([]int, int) {
	var i int = startat
	var sum []int
	var count int = 0

	for i <= endat {
		sum = append(sum, i)
		i++
		count++
	}
	return sum, count
}

func Avrfloat32(startat, endat float32) float32 {
	var i float32 = float32(startat)
	var sum float32 = 0
	var count int = 0
	for i <= float32(endat) {
		sum += i
		i++
		count++
	}
	return (sum / float32(count))
}
