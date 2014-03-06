package main

func printNum(i int) {
	if i == 0 {
		print('0')
	}

	base := 10
	for base < i {
		base *= 10
	}
	base /= 10

	for base > 0 {
		d := i / base
		i = i % base
		base /= 10
		print('0' + uint8(d))
	}
}

func main() {
	printNum(3927)
	print('\n')
}
