func printNum(i int) {
	if i == 0 {
		print('0')
	}

	var base = 10
	for base < i {
		base *= 10
	}
	base /= 10

	for base > 0 {
		var d = i / base
		i = i % base
		base /= 10
		print('0' + uint8(d))
	}
}

func main() {
	printNum(3927)
	print('\n')
}
