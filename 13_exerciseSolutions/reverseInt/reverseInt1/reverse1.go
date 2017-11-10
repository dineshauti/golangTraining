package reverseInt1

import (
	"strconv"
)

func rev(y []byte) int {

	for i:=0; i<len(y)/2; i++ {
		j := len(y) - i - 1
		y[i], y[j] = y[j], y[i]
	}
	data, _ := strconv.Atoi(string(y))
	return  data
}

func Reverse1(x int) int {

	y := []byte(strconv.Itoa(x))

	if y[0] == 45 {

		data := rev(y[1:]) * -1
		if data > -2147483648 {
			return data
		}

	} else {

		data := rev(y)
		if data < 2147483647 {
			return data
		}
	}
	return 0
}
