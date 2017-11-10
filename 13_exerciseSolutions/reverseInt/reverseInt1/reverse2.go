package reverseInt1


func Reverse2(x int) int{

	a := 0
	for x != 0 {
		a = a*10 + x%10
		x=x/10
	}

	if a > -2147483648 && a < 2147483647 {
		return a
	}

	return 0

}
