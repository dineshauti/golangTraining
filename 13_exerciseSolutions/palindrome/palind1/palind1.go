package palind1

func IsPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	originalValue := x
	reverseValue := 0
	for x != 0 {
		reverseValue = reverseValue*10 + x%10
		x=x/10
	}

	if originalValue == reverseValue {
		return true
	}
	return false
}
